package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) WriteJSON(w http.ResponseWriter, status int, msg, detail string, data any, headers ...http.Header) error {
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	payloadData := jsonResponse{
		Code:    status,
		Message: msg,
		Detail:  detail,
		Data:    data,
	}

	payload, err := json.Marshal(payloadData)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(payload)
	if err != nil {
		return err
	}

	return nil
}

func (app *Config) ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	return app.WriteJSON(w, statusCode, "error", err.Error(), nil)
}
