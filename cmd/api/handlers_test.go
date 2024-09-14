package main

import (
	"testing"
)

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("handler returned wrong status code: got %v want %v", got, want)
	}
}

func assertBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("handler returned unexpected body: got %v want %v", got, want)
	}
}

func assertContentType(t *testing.T, got string) {
	t.Helper()

	if got != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want application/json", got)
	}
}

// func TestGetTracks(t *testing.T) {
// 	tests := []struct {
// 		name           string
// 		url            string
// 		expectedStatus int
// 		expectedBody   string
// 	}{
// 		{
// 			name:           "Successful request",
// 			url:            "/getTracks",
// 			expectedStatus: http.StatusOK,
// 			expectedBody:   `{"name":"Name","year":2024}`,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			rr := httptest.NewRecorder()
// 			app := &Config{}
// 			handler := app.routes()

// 			req, err := http.NewRequest("GET", tt.url, nil)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			handler.ServeHTTP(rr, req)

// 			assertStatusCode(t, rr.Code, tt.expectedStatus)
// 			assertBody(t, rr.Body.String(), tt.expectedBody)
// 			assertContentType(t, rr.Header().Get("Content-Type"))
// 		})
// 	}
// }
