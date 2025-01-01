package project

import (
	"net/http"

	"ericarthurc.com/internal/view"
)

type handlers struct {
	*router
}

func newHandlers(router *router) *handlers {
	return &handlers{router}
}

// @Route: /
// @Method: GET
// @Render the project page
func (h *handlers) projectIndexHTML() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.TemplRender(w, r, 200, view.ProjectIndex(), true)
	}
}
