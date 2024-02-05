package webapi

import (
	"net/http"
	"strings"
)

type route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

// Router type.
type Router struct {
	routes []route
}

// NewRouter creates a new Router.
func NewRouter() *Router {
	rtr := &Router{
		routes: []route{},
	}

	return rtr
}

// HandleFunc registers a HandlerFunc for a given path and method.
func (r *Router) HandleFunc(method string, path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, route{method, path, handler})
}

// ServeHTTP handles the core routing logic.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Populated later with found valid methods for 405 responses.
	var allowed []string
	// Iterate over and check paths for match.
	for _, route := range r.routes {
		if route.path == req.URL.Path {
			if req.Method != route.method {
				// Wrong method. Add match to allowed methods for possible 405 response.
				allowed = append(allowed, route.method)
				continue
			}
			// Perfect match. Forward to handler.
			route.handler(w, req)
			return
		}
	}
	if len(allowed) > 0 {
		// Matching path(s) with a different http method requirement were found.
		w.Header().Set("Allow", strings.Join(allowed, ", "))
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// No matches were found at all.
	http.NotFound(w, req)
}
