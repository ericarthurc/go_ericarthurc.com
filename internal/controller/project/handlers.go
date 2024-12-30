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
		if err := h.TemplRender(w, r, 200, view.ProjectIndex()); err != nil {
			h.Error(w, http.StatusInternalServerError, "failed to render template")
		}
	}
}
