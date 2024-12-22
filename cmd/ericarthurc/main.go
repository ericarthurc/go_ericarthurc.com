package main

import (
	"log"
	"net/http"
	"os"

	"ericarthurc.com/internal/controllers/blog"
	"ericarthurc.com/internal/orbit"
	"ericarthurc.com/internal/views"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// initialize database pool
	// dbPool, err := database.NewDbPool()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	stylesRaw, err := os.ReadFile("web/compiled/css/main.css")
	if err != nil {
		log.Fatal("error loading compiled css file")
	}

	orb := orbit.NewOrbit(string(stylesRaw))
	r := chi.NewRouter()

	// root level middleware stack
	r.Use(chiMiddleware.RealIP)
	// r.Use(middleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	// serve static files
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/root/favicon.ico")
	})
	r.Get("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/root/robots.txt")
	})
	r.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir("web/static"))))

	// routes
	r.Group(func(r chi.Router) {

		// r.Mount("/")

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			orb.TemplRender(w, 200, views.IndexPage())
		})

		r.Mount("/blog", blog.Routes(orb))
	})

	orbit.Launch(r)
}
