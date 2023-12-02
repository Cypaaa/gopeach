package gopeach

import "github.com/valyala/fasthttp"

// RequestCtx is an extension of fasthttp.RequestCtx
type RequestCtx struct {
	*fasthttp.RequestCtx
	Params     map[string]string
	JsonBody   map[string]interface{}
	Nexts      []func(ctx *RequestCtx)
	pathString string
}

// NewRequestCtx returns a new RequestCtx
func NewRequestCtx(ctx *fasthttp.RequestCtx, h []func(ctx *RequestCtx)) *RequestCtx {
	//var x RequestCtx =
	return &RequestCtx{
		RequestCtx: ctx,
		Params:     make(map[string]string),
		JsonBody:   make(map[string]interface{}),
		Nexts:      h,
		pathString: string(ctx.Path()),
	}
}

// Next calls the next handler
func (ctx *RequestCtx) Next() {
	if len(ctx.Nexts) > 0 {
		h := ctx.Nexts[0]
		ctx.Nexts = ctx.Nexts[1:]
		h(ctx)
	}
}

// Send is and alias for WriteString
func (ctx *RequestCtx) Send(s string) {
	ctx.WriteString(s)
}

// Method returns the request method
func (ctx *RequestCtx) MethodString() string {
	return string(ctx.Method())
}

// Path returns the request path
func (ctx *RequestCtx) PathString() string {
	return ctx.pathString
}

// Cookie returns cookie value by name
func (ctx *RequestCtx) Cookie(name string) string {
	return string(ctx.Request.Header.Cookie(name))
}

// SetCookie sets cookie.
func (ctx *RequestCtx) SetCookie(cookie *fasthttp.Cookie) {
	ctx.Response.Header.SetCookie(cookie)
}

// GetHeader returns the request header with key
func (ctx *RequestCtx) Header(key string) string {
	return string(ctx.Request.Header.Peek(key))
}

// SetHeader sets the response header with key, value pair
func (ctx *RequestCtx) SetHeader(key, value string) {
	ctx.Response.Header.Set(key, value)
}

// SetContentType sets the response content type
func (ctx *RequestCtx) SetContentType(contentType string) {
	ctx.Response.Header.SetContentType(contentType)
}

// SetStatusCode sets the response status code
func (ctx *RequestCtx) SetStatusCode(statusCode int) {
	ctx.Response.SetStatusCode(statusCode)
}
