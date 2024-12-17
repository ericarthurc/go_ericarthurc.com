package blog

import (
	"ericarthurc.com/internal/orbit"
	"github.com/go-chi/chi/v5"
)

// Can pass global state to the router here
type blogRouter struct {
	*orbit.Orbit
	*chi.Mux
}

func newRouter(orb *orbit.Orbit) *blogRouter {
	return &blogRouter{
		Orbit: orb,
		Mux:   chi.NewRouter(),
	}
}

func Routes(orb *orbit.Orbit) *blogRouter {
	r := newRouter(orb)
	h := blogHandlers{r}

	r.Get("/", h.TemplRoot())

	return r
}
