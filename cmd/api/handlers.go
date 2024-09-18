package main

import (
	"fmt"
	"net/http"
	"strconv"

	_ "F1ResultsApi/data"

	"github.com/go-chi/chi/v5"
)

type APIHandler interface {
	HandleGetTracks(http.ResponseWriter, *http.Request)
	HandleGetTrack(http.ResponseWriter, *http.Request)
	HandleGetResults(http.ResponseWriter, *http.Request)
	HandleGetResult(http.ResponseWriter, *http.Request)
}

var _ APIHandler = (*Config)(nil)

// @Summary Get tracks for a specific year
// @Description Retrieves all tracks for the specified year
// @Tags tracks
// @Accept json
// @Produce json
// @Param year path int true "Year"
// @Success 200 {object} jsonResponse{Data=[]data.Track} "Successful operation"
// @Failure 400 {object} jsonResponse "Bad request"
// @Failure 404 {object} jsonResponse "Not found"
// @Router /getTracks/{year} [get]
func (app *Config) HandleGetTracks(w http.ResponseWriter, r *http.Request) {
	yearParam := chi.URLParam(r, "year")
	year, err := strconv.Atoi(yearParam)
	if err != nil {
		app.ErrorJSON(w, fmt.Errorf("invalid parameter YEAR"), http.StatusBadRequest)
		return
	}

	data, err := app.Repository.GetTracks(year)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if len(*data) == 0 {
		app.WriteJSON(w, http.StatusNotFound, "error", "not found", nil)
		return
	}

	app.WriteJSON(w, http.StatusOK, "success", "", data)
}

// @Summary Get track for a specific year and track name
// @Description Retrieves specific track for the specified year and track name
// @Tags tracks
// @Accept json
// @Produce json
// @Param year path int true "Year"
// @Param trackName path string true "Track Name"
// @Success 200 {object} jsonResponse{Data=data.Track} "Successful operation"
// @Failure 400 {object} jsonResponse "Bad request"
// @Failure 404 {object} jsonResponse "Not found"
// @Router /getTracks/{year}/{trackName} [get]
func (app *Config) HandleGetTrack(w http.ResponseWriter, r *http.Request) {
	trackName := chi.URLParam(r, "trackName")
	yearParam := chi.URLParam(r, "year")
	year, err := strconv.Atoi(yearParam)
	if err != nil {
		app.ErrorJSON(w, fmt.Errorf("invalid parameter YEAR"), http.StatusBadRequest)
		return
	}

	data, err := app.Repository.GetTrack(year, trackName)
	if err != nil {
		switch err.Error() {
		case "sql: no rows in result set":
			app.WriteJSON(w, http.StatusNotFound, "error", "not found", nil)
		default:
			app.ErrorJSON(w, err, http.StatusBadRequest)
		}

		return
	}

	if data == nil {
		app.WriteJSON(w, http.StatusNotFound, "error", "not found", nil)
		return
	}

	app.WriteJSON(w, http.StatusOK, "success", "", data)
}

// @Summary Get results for a specific year
// @Description Retrieves specific results for the specified year
// @Tags results
// @Accept json
// @Produce json
// @Param year path int true "Year"
// @Success 200 {object} jsonResponse{Data=[]data.Result} "Successful operation"
// @Failure 400 {object} jsonResponse "Bad request"
// @Failure 404 {object} jsonResponse "Not found"
// @Router /getResults/{year} [get]
func (app *Config) HandleGetResults(w http.ResponseWriter, r *http.Request) {
	yearParam := chi.URLParam(r, "year")
	year, err := strconv.Atoi(yearParam)
	if err != nil {
		app.ErrorJSON(w, fmt.Errorf("invalid parameter YEAR"), http.StatusBadRequest)
		return
	}

	data, err := app.Repository.GetResults(year)
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if len(*data) == 0 {
		app.WriteJSON(w, http.StatusNotFound, "error", "not found", nil)
		return
	}

	app.WriteJSON(w, http.StatusOK, "success", "", data)
}

// @Summary Get results for a specific year and track ID
// @Description Retrieves specific results for the specified year and track ID
// @Tags results
// @Accept json
// @Produce json
// @Param year path int true "Year"
// @Param trackId path int true "Track ID"
// @Success 200 {object} jsonResponse{Data=[]data.Result} "Successful operation"
// @Failure 400 {object} jsonResponse "Bad request"
// @Failure 404 {object} jsonResponse "Not found"
// @Router /getResult/{year}/{trackId} [get]
func (app *Config) HandleGetResult(w http.ResponseWriter, r *http.Request) {
	yearParam := chi.URLParam(r, "year")
	year, err := strconv.Atoi(yearParam)
	if err != nil {
		app.ErrorJSON(w, fmt.Errorf("invalid parameter YEAR"), http.StatusBadRequest)
		return
	}

	trackIdParam := chi.URLParam(r, "trackId")
	trackId, err := strconv.Atoi(trackIdParam)
	if err != nil {
		app.ErrorJSON(w, fmt.Errorf("invalid parameter TRACK_ID"), http.StatusBadRequest)
		return
	}

	data, err := app.Repository.GetResult(year, int64(trackId))
	if err != nil {
		switch err.Error() {
		case "sql: no rows in result set":
			app.WriteJSON(w, http.StatusNotFound, "error", "not found", nil)
		default:
			app.ErrorJSON(w, err, http.StatusBadRequest)
		}

		return
	}

	if len(*data) == 0 {
		app.WriteJSON(w, http.StatusNotFound, "error", "not found", nil)
		return
	}

	app.WriteJSON(w, http.StatusOK, "success", "", data)
}
