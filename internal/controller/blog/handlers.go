package blog

import (
	"net/http"

	"ericarthurc.com/internal/view"
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
// @Render the blog index page
func (h *handlers) blogIndexHTML() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.State.PostMeta.Mu.RLock()
		defer h.State.PostMeta.Mu.RUnlock()

		h.TemplRender(w, r, 200, view.BlogIndex(h.State.PostMeta.FeaturedPostsMetaSorted, h.State.PostMeta.NonFeaturedPostsMetaSorted), true)
	}
}

// @Route: /blog/{slug}
// @Method: GET
// @Render the blog post page dynamically by slug
func (h *handlers) blogSlugHTML() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")

		post, _ := h.State.PostMap.Load(slug)

		h.TemplRender(w, r, 200, view.BlogSlug(post), false)
	}
}
