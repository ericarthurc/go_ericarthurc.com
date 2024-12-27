package index

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
// @Render the index page
func (h *handlers) indexHTML() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h.TemplRender(w, 200, view.Index(h.State.PostMeta.FeaturedPostsMetaSorted, h.State.PostMeta.NonFeaturedPostsMetaSorted)); err != nil {
			h.Error(w, http.StatusInternalServerError, "failed to render template")
		}
	}
}
