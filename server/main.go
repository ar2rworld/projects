package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var SERVER_PORT = os.Getenv("SERVER_PORT")

func main() {
	SERVER_PORT = "8050"

	r := chi.NewRouter()

	kc, err := newKeycloak()
	if err != nil {
		log.Fatal(err)
	}

	controller := newController(kc)

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://ar2rworld-projects.vercel.app", "http://localhost:5173"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))

	r.Route("/projectsapi", func(r chi.Router) {

		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			// no auth
			if w.Header().Get("Authorization") == "" {
 				controller.login(w, r)
			}
		})
		
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("get /"))
		})
	})

	log.Printf("Server started on :%s", SERVER_PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", SERVER_PORT), r)
}
