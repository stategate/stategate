package auth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/autom8ter/stategate/internal/helpers"
	"github.com/autom8ter/stategate/internal/logger"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/open-policy-agent/opa/rego"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"sync"
)

type ctxKey string

var (
	userCtxKey ctxKey = "user-ctx-key"
)

const (
	reqDenied  = "request authorization = denied"
	respDenied = "response authorization = denied"
)

type Auth struct {
	jwksUri        string
	jwksSet        *jwk.Set
	mu             sync.RWMutex
	logger         *logger.Logger
	requestPolicy  rego.PreparedEvalQuery
	responsePolicy rego.PreparedEvalQuery
}

func NewAuth(jwksUri string, logger2 *logger.Logger, reqPolicy, respPolicy *rego.Rego) (*Auth, error) {
	respeval, err := respPolicy.PrepareForEval(context.Background())
	if err != nil {
		return nil, err
	}
	reqeval, err := reqPolicy.PrepareForEval(context.Background())
	if err != nil {
		return nil, err
	}

	a := &Auth{
		jwksUri:        jwksUri,
		jwksSet:        nil,
		mu:             sync.RWMutex{},
		logger:         logger2,
		requestPolicy:  reqeval,
		responsePolicy: respeval,
	}
	return a, a.RefreshJWKS()
}

func (a *Auth) RefreshJWKS() error {
	if a.jwksUri != "" {
		jwks, err := jwk.Fetch(a.jwksUri)
		if err != nil {
			return err
		}
		a.mu.Lock()
		a.jwksSet = jwks
		a.mu.Unlock()
	}

	return nil
}

func (a *Auth) ParseAndVerify(token string) ([]byte, error) {
	message, err := jws.ParseString(token)
	if err != nil {
		return nil, err
	}
	var payload []byte
	if a.jwksSet != nil && a.jwksUri != "" {
		a.mu.RLock()
		defer a.mu.RUnlock()
		if len(message.Signatures()) == 0 {
			return nil, fmt.Errorf("zero jws signatures")
		}
		kid, ok := message.Signatures()[0].ProtectedHeaders().Get("kid")
		if !ok {
			return nil, fmt.Errorf("jws kid not found")
		}
		algI, ok := message.Signatures()[0].ProtectedHeaders().Get("alg")
		if !ok {
			return nil, fmt.Errorf("jw alg not found")
		}
		alg, ok := algI.(jwa.SignatureAlgorithm)
		if !ok {
			return nil, fmt.Errorf("alg type cast error")
		}

		keys := a.jwksSet.LookupKeyID(kid.(string))
		if len(keys) == 0 {
			return nil, errors.Errorf("failed to lookup kid: %s - zero keys", kid.(string))
		}
		var key interface{}
		if err := keys[0].Raw(&key); err != nil {
			return nil, err
		}
		payload, err = jws.Verify([]byte(token), alg, key)
		if err != nil {
			return nil, err
		}
	} else {
		payload = message.Payload()
	}
	return payload, nil
}

func (a *Auth) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
		if err != nil {
			return nil, err
		}
		payload, err := a.ParseAndVerify(token)
		if err != nil {
			a.logger.Error(err.Error())
			return nil, status.Error(codes.Unauthenticated, "unverified")
		}
		claims := map[string]interface{}{}
		if err := json.Unmarshal(payload, &claims); err != nil {
			a.logger.Error(err.Error())
			return nil, status.Error(codes.Internal, "failed to parse claims")
		}
		md := metautils.ExtractIncoming(ctx)
		c := &Context{
			authPayload: base64.StdEncoding.EncodeToString(payload),
			Claims:      claims,
			Method:      info.FullMethod,
			Metadata:    map[string]string{},
			Body:        toMap(req),
		}
		for k, arr := range md {
			c.Metadata[k] = arr[0]
		}
		allowed, err := a.evaluateRequest(ctx, c)
		if err != nil {
			a.logger.Error(err.Error())
			return nil, status.Error(codes.Internal, "failed to evaluate authz policy")
		}
		if !allowed {
			return nil, status.Error(codes.PermissionDenied, reqDenied)
		}
		ctx = SetContext(ctx, c)
		hresp, err := handler(ctx, req)
		if err != nil {
			return resp, err
		}
		md = metautils.ExtractOutgoing(ctx)
		respMeta := make(map[string]string)
		for k, arr := range md {
			respMeta[k] = arr[0]
		}
		c.Metadata = respMeta
		allowed, err = a.evaluateResponse(ctx, c)
		if err != nil {
			a.logger.Error(err.Error())
			return nil, status.Error(codes.Internal, "failed to evaluate authz policy")
		}
		if !allowed {
			return nil, status.Error(codes.PermissionDenied, respDenied)
		}
		return hresp, nil
	}
}

func (a *Auth) StreamInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		token, err := grpc_auth.AuthFromMD(ss.Context(), "Bearer")
		if err != nil {
			return err
		}
		payload, err := a.ParseAndVerify(token)
		if err != nil {
			a.logger.Error(err.Error())
			return status.Error(codes.Unauthenticated, "unverified")
		}
		claims := map[string]interface{}{}
		if err := json.Unmarshal(payload, &claims); err != nil {
			a.logger.Error(err.Error())
			return status.Error(codes.Internal, "failed to parse claims")
		}
		md := metautils.ExtractIncoming(ss.Context())
		c := &Context{
			authPayload:  base64.StdEncoding.EncodeToString(payload),
			Claims:       claims,
			Method:       info.FullMethod,
			Metadata:     map[string]string{},
			Body:         map[string]interface{}{},
			ClientStream: info.IsClientStream,
			ServerStream: info.IsServerStream,
		}
		for k, arr := range md {
			if len(arr) > 0 {
				c.Metadata[k] = arr[0]
			}
		}
		ctx := SetContext(ss.Context(), c)
		return handler(srv, &stream{
			ss:  ss,
			a:   a,
			ctx: ctx,
		})
	}
}

func (a *Auth) evaluateRequest(ctx context.Context, context *Context) (bool, error) {
	results, err := a.requestPolicy.Eval(ctx, rego.EvalInput(context.input()))
	if err != nil {
		return false, errors.Wrap(err, "request policy: failed to evaluate input")
	}
	if len(results) == 0 {
		return false, errors.Wrap(err, "request policy: zero results")
	}
	if len(results[0].Expressions) == 0 {
		return false, errors.Wrap(err, "request policy: zero result expressions")
	}
	if results[0].Expressions[0].Value == nil {
		return false, errors.Wrap(err, "request policy: empty expression value")
	}
	res, ok := results[0].Expressions[0].Value.(bool)
	if !ok {
		return false, errors.Wrap(err, "request policy: expression does not return a boolean value")
	}
	return res, nil
}

func (a *Auth) evaluateResponse(ctx context.Context, context *Context) (bool, error) {
	results, err := a.responsePolicy.Eval(ctx, rego.EvalInput(context.input()))
	if err != nil {
		return false, errors.Wrap(err, "response policy: failed to evaluate input")
	}
	if len(results) == 0 {
		return false, errors.Wrap(err, "response policy: zero results")
	}
	if len(results[0].Expressions) == 0 {
		return false, errors.Wrap(err, "response policy: zero result expressions")
	}
	if results[0].Expressions[0].Value == nil {
		return false, errors.Wrap(err, "response policy: empty expression value")
	}
	res, ok := results[0].Expressions[0].Value.(bool)
	if !ok {
		return false, errors.Wrap(err, "response policy: expression does not return a boolean value")
	}
	return res, nil
}

func toMap(obj interface{}) map[string]interface{} {
	data := map[string]interface{}{}
	if val, ok := obj.(proto.Message); ok {
		bits, _ := helpers.MarshalJSON(val)
		json.Unmarshal(bits, &data)
	} else {
		bits, _ := json.Marshal(obj)
		json.Unmarshal(bits, &data)
	}
	return data
}
