package blog

import (
	"ericarthurc.com/internal/controller/state"
	"ericarthurc.com/internal/orbit"
	"github.com/go-chi/chi/v5"
)

// Can pass global state to the router here
type router struct {
	*state.State
	*orbit.Orbit
	*chi.Mux
}

func newRouter(st *state.State, orb *orbit.Orbit) *router {
	return &router{
		State: st,
		Orbit: orb,
		Mux:   chi.NewRouter(),
	}
}

// Mounted: /blog
func Routes(st *state.State, orb *orbit.Orbit) *router {
	r := newRouter(st, orb)
	h := newHandlers(r)

	r.Get("/", h.blogIndexHTML())
	r.Get("/{slug}", h.blogSlugHTML())

	return r
}
