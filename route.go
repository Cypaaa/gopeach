package gopeach

import "strings"

// Route stores path, param indexes and handler.
type Route struct {
	Path         string
	ParamIndexes map[string]uint // map[string]uint{"id": 0, "name": 1, ...}
	Handler      func(ctx *RequestCtx)
}

// NewRoute takes a path string and an handler function and returns new Route.
func NewRoute(path string, h func(ctx *RequestCtx)) Route {
	r := convertPath(path)
	r.Handler = h
	return r
}

// convertPath converts path to regexp and returns Route struct.
//
// e.g. "/users/:id" -> "/users/[a-zA-Z0-9-@:%._\\+~#?&=]{1,256}"
func convertPath(s string) Route {
	pi := make(map[string]uint)
	ps := strings.Split(s, "/")
	for i, p := range ps {
		if strings.HasPrefix(p, ":") {
			pi[p[1:]] = uint(i)
			ps[i] = `[a-zA-Z0-9-@:%._\\+~#?&=]{1,256}`
		}
	}
	ps[0] = "^" + ps[0]
	ps[len(ps)-1] = ps[len(ps)-1] + "$"

	return Route{
		Path:         strings.Join(ps, "/"),
		ParamIndexes: pi,
		Handler:      nil,
	}
}
