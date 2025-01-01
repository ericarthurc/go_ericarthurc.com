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
		h.State.PostMeta.Mu.RLock()
		defer h.State.PostMeta.Mu.RUnlock()

		h.TemplRender(w, r, 200, view.Index(h.State.PostMeta.FeaturedPostsMetaSorted, h.State.PostMeta.NonFeaturedPostsMetaSorted), true)
	}
}
