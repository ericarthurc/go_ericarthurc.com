package index

import (
	"ericarthurc.com/internal/orbit"
	"github.com/go-chi/chi/v5"
)

// Can pass global state to the router here
type indexRouter struct {
	*orbit.Orbit
	*chi.Mux
}

func newRouter(orb *orbit.Orbit) *indexRouter {
	return &indexRouter{
		Orbit: orb,
		Mux:   chi.NewRouter(),
	}
}

// route: /
func Routes(orb *orbit.Orbit) *indexRouter {
	r := newRouter(orb)
	// h := indexHandlers{r}

	// r.Get("/", h.TemplRoot())

	return r
}
