package webapi

import (
	"github.com/bhwz/weather/weather-server/webapi/services"
	"net/http/pprof"
)

// All global routes are defined here.
func (s *server) setRoutes() {
	s.router.HandleFunc("GET", "/", services.StatusService())

	if s.config.Debug {
		s.router.HandleFunc("GET", "/debug/pprof/", pprof.Index)
		s.router.HandleFunc("GET", "/debug/pprof/cmdline", pprof.Cmdline)
		s.router.HandleFunc("GET", "/debug/pprof/profile", pprof.Profile)
		s.router.HandleFunc("GET", "/debug/pprof/symbol", pprof.Symbol)
		s.router.HandleFunc("POST", "/debug/pprof/symbol", pprof.Symbol)
		s.router.HandleFunc("GET", "/debug/pprof/trace", pprof.Trace)
		s.router.HandleFunc("GET", "/debug/pprof/allocs", pprof.Handler("allocs").ServeHTTP)
		s.router.HandleFunc("GET", "/debug/pprof/block", pprof.Handler("block").ServeHTTP)
		s.router.HandleFunc("GET", "/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)
		s.router.HandleFunc("GET", "/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
		s.router.HandleFunc("GET", "/debug/pprof/mutex", pprof.Handler("mutex").ServeHTTP)
		s.router.HandleFunc("GET", "/debug/pprof/threadcreate", pprof.Handler("threadcreate").ServeHTTP)
	}
}
