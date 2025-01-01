package orbit

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"ericarthurc.com/internal/view"
	"github.com/a-h/templ"
)

type Orbit struct {
	GlobalStyles string
}

func NewOrbit(styles string) *Orbit {
	return &Orbit{
		GlobalStyles: styles,
	}
}

func Launch(r http.Handler) {
	fmt.Printf("🔥 Launching at http://%s:%s 🔥\n", os.Getenv("HOST"), os.Getenv("PORT"))
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")), r)
	if err != nil {
		log.Fatal(err)
	}
}

func (o *Orbit) Text(w http.ResponseWriter, code int, text string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	w.Write([]byte(text))
}

func (o *Orbit) HTML(w http.ResponseWriter, code int, html string) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(code)
	w.Write([]byte(html))
}

func (o *Orbit) TemplRender(w http.ResponseWriter, r *http.Request, code int, component templ.Component, cache bool) {
	if cache {
		w.Header().Set("Cache-Control", "private, max-age=60")
	}

	// wrap the temple view in the main layout
	view.Main(o.GlobalStyles, component, r.URL.Path).Render(context.Background(), w)

}

func (o *Orbit) Error(w http.ResponseWriter, code int, errorMessage string) {
	http.Error(w, errorMessage, code)
}
