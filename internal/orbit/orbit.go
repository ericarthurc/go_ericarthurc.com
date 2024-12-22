package orbit

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"ericarthurc.com/internal/views"
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

func (o *Orbit) TemplRender(w http.ResponseWriter, code int, view templ.Component) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(code)
	if err := views.Main(o.GlobalStyles, view).Render(context.Background(), w); err != nil {
		return err
	}

	return nil
}

func (o *Orbit) Error(w http.ResponseWriter, code int, errorMessage string) {
	http.Error(w, errorMessage, code)
}
