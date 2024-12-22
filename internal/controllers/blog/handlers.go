package blog

import (
	"net/http"

	"ericarthurc.com/internal/views"
)

type blogHandlers struct {
	*blogRouter
}

func newBlogHandlers(router *blogRouter) *blogHandlers {
	return &blogHandlers{router}
}

func (h *blogHandlers) TemplRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h.TemplRender(w, 200, views.BlogPage("John")); err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
		}
	}
}
