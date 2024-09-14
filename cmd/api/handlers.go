package main

import (
	"encoding/json"
	"net/http"
)

type APIHandler interface {
	HandleGetTracks(http.ResponseWriter, *http.Request)
	HandleGetTrack(http.ResponseWriter, *http.Request)
	HandleGetResults(http.ResponseWriter, *http.Request)
	HandleGetResult(http.ResponseWriter, *http.Request)
}

var _ APIHandler = (*Config)(nil)

func (app *Config) HandleGetTracks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	d, _ := app.Repository.GetTracks(2024)
	dm, _ := json.Marshal(d)

	_, err := w.Write(dm)
	if err != nil {
		panic(err)
	}
}

func (app *Config) HandleGetTrack(w http.ResponseWriter, r *http.Request) {}

func (app *Config) HandleGetResults(w http.ResponseWriter, r *http.Request) {}

func (app *Config) HandleGetResult(w http.ResponseWriter, r *http.Request) {}
