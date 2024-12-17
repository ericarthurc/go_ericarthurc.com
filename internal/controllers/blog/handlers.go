package blog

import (
	"net/http"

	"ericarthurc.com/internal/views"
)

type blogHandlers struct {
	*blogRouter
}

func (h *blogHandlers) TemplRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.TemplRender(w, 200, views.BlogPage("John"))
	}
}
