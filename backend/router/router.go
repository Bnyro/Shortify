package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/url-shortener/entities"
	"github.com/url-shortener/handlers"
)

func Create() {
	router := chi.NewRouter()

	router.Use(cors.AllowAll().Handler)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(
			entities.Message{
				Message: "API online",
			},
		)
	})

	router.Post("/create", handlers.CreateShortcut)
	router.Get("/{short}", handlers.ReadShortcut)

	log.Fatal(http.ListenAndServe(":5000", router))
}
