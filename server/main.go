package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/ar2rworld/projects-backend/controller"
)

var SERVER_PORT = os.Getenv("SERVER_PORT")

func main() {
	SERVER_PORT = "8050"

	r := chi.NewRouter()

	controller := controller.NewController()
	controller.NotificatorToken = os.Getenv("NOTIFICATOR_TOKEN")
	controller.NotificatorChatID = os.Getenv("NOTIFICATOR_CHAT_ID")

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://ar2rworld-projects.vercel.app", "http://localhost:5173"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))

	r.Route("/api", func(r chi.Router) {

		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			// no auth
			if w.Header().Get("Authorization") == "" {
 				controller.KratosLogin(w, r)
			}
		})

		r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
			// no auth
			if w.Header().Get("Authorization") == "" {
 				controller.Register(w, r)
			}
		})
		
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("get /"))
		})

		r.Get("/me", func (w http.ResponseWriter, r *http.Request) {
			controller.GetMe(w, r)
		})
		r.Post("/contact", func(w http.ResponseWriter, r *http.Request) {
			controller.Contact(w, r)
		})
	})

	log.Printf("Server started on :%s", SERVER_PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", SERVER_PORT), r)
}
