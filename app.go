package gopeach

import (
	"regexp"
	"strings"

	"github.com/valyala/fasthttp"
)

// match matches path, stores user values and calls handler
//
// e.g. "/users/1" -> "/users/:id" where "id" is stored as key and "1" as value in user values
func match(ctx *RequestCtx, r []Route) func(ctx *RequestCtx) {
	for _, route := range r {
		ps := strings.Split(string(ctx.Path()), "/")
		if ok, err := regexp.Match(route.Path, ctx.Path()); ok && err == nil {
			for k, v := range route.ParamIndexes {
				ctx.Params[k] = ps[v]
			}
			return route.Handler
		}
	}
	return func(ctx *RequestCtx) {
		ctx.Error("Not Found", fasthttp.StatusNotFound)
	}
}

// App stores Routes and middlewares
type App struct {
	get           []Route
	post          []Route
	patch         []Route
	put           []Route
	delete        []Route
	middleware    []func(ctx *RequestCtx)
	caseSensitive bool // Default is true
}

// New returns a new App
func New() *App {
	return &App{
		get:           []Route{},
		post:          []Route{},
		patch:         []Route{},
		put:           []Route{},
		delete:        []Route{},
		middleware:    []func(ctx *RequestCtx){},
		caseSensitive: true,
	}
}

// CaseSensitive sets if the routes should be checked with case sensitivity
func (r *App) CaseSensitive(b bool) {
	r.caseSensitive = b
}

// Listen listens on addr
func (r *App) Listen(addr string) error {
	return fasthttp.ListenAndServe(addr, r.Handler)
}

// ListenTLS listens on addr with certFile and keyFile
func (r *App) ListenTLS(addr, certFile, keyFile string) error {
	return fasthttp.ListenAndServeTLS(addr, certFile, keyFile, r.Handler)
}

// Get adds GET Route
func (r *App) Get(s string, h func(ctx *RequestCtx)) {
	r.get = append(r.get, NewRoute(s, h))
}

// Post adds POST Route
func (r *App) Post(s string, h func(ctx *RequestCtx)) {
	r.post = append(r.post, NewRoute(s, h))
}

// Patch adds PATCH Route
func (r *App) Patch(s string, h func(ctx *RequestCtx)) {
	r.patch = append(r.patch, NewRoute(s, h))
}

// Put adds PUT Route
func (r *App) Put(s string, h func(ctx *RequestCtx)) {
	r.put = append(r.put, NewRoute(s, h))
}

// Delete adds DELETE Route.
func (r *App) Delete(s string, h func(ctx *RequestCtx)) {
	r.delete = append(r.delete, NewRoute(s, h))
}

// Middleware adds middleware
func (r *App) Middleware(h func(ctx *RequestCtx)) {
	r.middleware = append(r.middleware, h)
}

// Handler handles requests
func (r *App) Handler(ctx *fasthttp.RequestCtx) {
	var h func(ctx *RequestCtx)
	req := NewRequestCtx(ctx, r.middleware)

	switch string(ctx.Method()) {
	case "GET":
		h = match(req, r.get)
	case "POST":
		h = match(req, r.post)
	case "PATCH":
		h = match(req, r.patch)
	case "PUT":
		h = match(req, r.put)
	case "DELETE":
		h = match(req, r.delete)
	}

	if h != nil {
		req.Nexts = append(req.Nexts, h)
	}
	req.Next()
}
