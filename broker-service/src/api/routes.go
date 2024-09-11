package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // NOTE: allow all hosts
		AllowedMethods:   []string{"POST", "GET", "DELETE", "PATCH", "OPTIONS", "PUT"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "x-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/", app.Root)

	r.Post("/", app.Broker)

	return r

}
