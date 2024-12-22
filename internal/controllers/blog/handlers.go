package blog

import (
	"net/http"

	"ericarthurc.com/internal/views"
	"github.com/go-chi/chi/v5"
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
		if err := h.TemplRender(w, 200, views.BlogPage()); err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
		}
	}
}

// @Route: /blog/{slug}
// @Method: GET
// @Render the blog post
func (h *handlers) blogSlugHTML() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")

		h.Text(w, 200, "Blog post: "+slug)
	}
}
