package auth

import "context"

type Context struct {
	Claims       map[string]interface{} `json:"claims"`
	Method       string                 `json:"method"`
	Request      map[string]interface{} `json:"request"`
	Response     map[string]interface{} `json:"response"`
	Headers      map[string]string      `json:"headers"`
	ClientStream bool                   `json:"client_stream"`
	ServerStream bool                   `json:"server_stream"`
}

func (a *Context) input() map[string]interface{} {
	return map[string]interface{}{
		"claims":        a.Claims,
		"method":        a.Method,
		"headers":       a.Headers,
		"request":       a.Request,
		"response":      a.Response,
		"client_stream": a.ClientStream,
		"server_stream": a.ClientStream,
	}
}

func SetContext(ctx context.Context, contxt *Context) context.Context {
	return context.WithValue(ctx, userCtxKey, contxt)
}

func GetContext(ctx context.Context) (*Context, bool) {
	if ctx.Value(userCtxKey) == nil {
		return nil, false
	}
	data, ok := ctx.Value(userCtxKey).(*Context)
	return data, ok
}
