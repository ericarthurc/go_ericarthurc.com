package blog

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

// @Route: /blog
// @Method: GET
// @Render the blog index
func (h *handlers) blogIndexHTML() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h.TemplRender(w, 200, view.BlogIndex()); err != nil {
			h.Error(w, http.StatusInternalServerError, "failed to render template")
		}
	}
}

// @Route: /blog/{slug}
// @Method: GET
// @Render the blog post
func (h *handlers) blogSlugHTML() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// slug := chi.URLParam(r, "slug")

		h.TemplRender(w, 200, view.BlogSlug())
	}
}
