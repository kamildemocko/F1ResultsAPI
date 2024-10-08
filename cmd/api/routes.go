package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "F1ResultsApi/docs"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Accept", "Content-Type", "X-CSRD-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(logMiddleware)

	mux.Route("/f1/api", func(mux chi.Router) {
		mux.Get("/getTracks/{year}", app.HandleGetTracks)
		mux.Get("/getTrack/{year}/{trackName}", app.HandleGetTrack)
		mux.Get("/getResults/{year}", app.HandleGetResults)
		mux.Get("/getResult/{year}/{trackId}", app.HandleGetResult)

		mux.Get("/swagger/*", httpSwagger.WrapHandler)
	})

	return mux
}
