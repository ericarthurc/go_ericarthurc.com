package about

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
// @Render the about page
func (h *handlers) aboutIndexHTML() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h.TemplRender(w, r, 200, view.AboutIndex()); err != nil {
			h.Error(w, http.StatusInternalServerError, "failed to render template")
		}
	}
}
