package auth

import (
	"context"
)

type Context struct {
	Claims       map[string]interface{} `json:"claims"`
	Method       string                 `json:"method"`
	Metadata     map[string]string      `json:"metadata"`
	Body         map[string]interface{} `json:"body"`
	ClientStream bool                   `json:"client_stream"`
	ServerStream bool                   `json:"server_stream"`
}

func (a *Context) input() map[string]interface{} {
	return map[string]interface{}{
		"claims":        a.Claims,
		"method":        a.Method,
		"metadata":      a.Metadata,
		"body":          a.Body,
		"client_stream": a.ClientStream,
		"server_stream": a.ClientStream,
	}
}

func (c *Context) copy() *Context {
	return &Context{
		Claims:       c.Claims,
		Method:       c.Method,
		Body:         c.Body,
		Metadata:     c.Metadata,
		ClientStream: c.ClientStream,
		ServerStream: c.ServerStream,
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
	if ok {
		return data.copy(), true
	}
	return data, ok
}
