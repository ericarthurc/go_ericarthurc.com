package main

import (
	"log"
	"net/http"
	"os"

	"ericarthurc.com/internal/controller/blog"
	"ericarthurc.com/internal/controller/index"
	"ericarthurc.com/internal/database"
	"ericarthurc.com/internal/orbit"
	"ericarthurc.com/internal/state"
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
	dbPool, err := database.NewDbPool()
	if err != nil {
		log.Fatal(err)
	}

	state, err := state.NewState(dbPool)
	if err != nil {
		log.Fatal(err)
	}

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
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// routes
	r.Group(func(r chi.Router) {

		r.Mount("/", index.Routes(state, orb))
		r.Mount("/blog", blog.Routes(state, orb))
	})

	orbit.Launch(r)
}
