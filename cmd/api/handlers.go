package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type APIHandler interface {
	HandleGetTracks(http.ResponseWriter, *http.Request)
	HandleGetTrack(http.ResponseWriter, *http.Request)
	HandleGetResults(http.ResponseWriter, *http.Request)
	HandleGetResult(http.ResponseWriter, *http.Request)
}

var _ APIHandler = (*Config)(nil)

func (app *Config) HandleGetTracks(w http.ResponseWriter, r *http.Request) {
	yearPram := chi.URLParam(r, "year")
	year, err := strconv.Atoi(yearPram)
	if err != nil {
		app.ErrorJSON(w, fmt.Errorf("invalid parameter YEAR"), http.StatusBadRequest)
		return
	}

	data, err := app.Repository.GetTracks(year)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if len(data) == 0 {
		app.WriteJSON(w, http.StatusNotFound, "error", "not found", nil)
		return
	}

	app.WriteJSON(w, http.StatusOK, "success", "", data)

}

func (app *Config) HandleGetTrack(w http.ResponseWriter, r *http.Request) {}

func (app *Config) HandleGetResults(w http.ResponseWriter, r *http.Request) {}

func (app *Config) HandleGetResult(w http.ResponseWriter, r *http.Request) {}
