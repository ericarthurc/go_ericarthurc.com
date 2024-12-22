package blog

import (
	"ericarthurc.com/internal/orbit"
	"github.com/go-chi/chi/v5"
)

// Can pass global state to the router here
type router struct {
	*orbit.Orbit
	*chi.Mux
}

func newRouter(orb *orbit.Orbit) *router {
	return &router{
		Orbit: orb,
		Mux:   chi.NewRouter(),
	}
}

// Mounted: /blog
func Routes(orb *orbit.Orbit) *router {
	r := newRouter(orb)
	h := newHandlers(r)

	r.Get("/", h.blogIndexHTML())
	r.Get("/{slug}", h.blogSlugHTML())

	return r
}
