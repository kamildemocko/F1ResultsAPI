package main

import (
	"F1ResultsApi/data"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
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

func assertContains(t *testing.T, got, want_substring string) {
	t.Helper()

	if !strings.Contains(got, want_substring) {
		t.Errorf("handler returned unexpected error: got %v want %v", got, want_substring)
	}
}

func TestDBProblems(t *testing.T) {
	tests := []struct {
		name                   string
		url                    string
		dsn                    string
		expected_err_substring string
	}{
		{
			name:                   "Bad DSN host",
			url:                    "/getTracks/2024",
			dsn:                    `host=192.168.0.244 port=5432 user=postgres password=badpassword sslmode=disable timezone=UTC connect_timeout=5 search_path=f1scrap`,
			expected_err_substring: "failed to connect",
		},
		{
			name:                   "Bad DSN password",
			url:                    "/getTracks/2024",
			dsn:                    `host=192.168.0.10 port=5432 user=postgres password=badpassword sslmode=disable timezone=UTC connect_timeout=5 search_path=f1scrap`,
			expected_err_substring: "password authentication failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := initPostgresDB(tt.dsn)

			assertContains(t, err.Error(), tt.expected_err_substring)
		})
	}
}

func TestGetTracks(t *testing.T) {
	godotenv.Load(".test.env")
	dsn := os.Getenv("DSN")

	db, err := initPostgresDB(dsn)
	if err != nil {
		panic(err)
	}

	repo := data.NewRepository(db)

	tests := []struct {
		name           string
		url            string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Successful request",
			url:            "/getTracks/2023",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"code":200,"message":"success","detail":"","data":[{"id":23,"name":"GULF AIR BAHRAIN GRAND PRIX","link":"https://www.formula1.com/en/results/2023/races/1141/bahrain/race-result","year":2023},{"id":24,"name":"STC SAUDI ARABIAN GRAND PRIX","link":"https://www.formula1.com/en/results/2023/races/1142/saudi-arabia/race-result","year":2023},{"id":25,"name":"ROLEX AUSTRALIAN GRAND PRIX","link":"https://www.formula1.com/en/results/2023/races/1143/australia/race-result","year":2023},{"id":26,"name":"AZERBAIJAN GRAND PRIX","link":"https://www.formula1.com/en/results/2023/races/1207/azerbaijan/race-result","year":2023},{"id":27,"name":"CRYPTO.COM MIAMI GRAND PRIX","link":"https://www.formula1.com/en/results/2023/races/1208/miami/race-result","year":2023},{"id":28,"name":"GRAND PRIX DE MONACO","link":"https://www.formula1.com/en/results/2023/races/1210/monaco/race-result","year":2023},{"id":29,"name":"AWS GRAN PREMIO DE ESPAÑA","link":"https://www.formula1.com/en/results/2023/races/1211/spain/race-result","year":2023},{"id":30,"name":"PIRELLI GRAND PRIX DU CANADA","link":"https://www.formula1.com/en/results/2023/races/1212/canada/race-result","year":2023},{"id":31,"name":"ROLEX GROSSER PREIS VON ÖSTERREICH","link":"https://www.formula1.com/en/results/2023/races/1213/austria/race-result","year":2023},{"id":32,"name":"ARAMCO BRITISH GRAND PRIX","link":"https://www.formula1.com/en/results/2023/races/1214/great-britain/race-result","year":2023},{"id":33,"name":"QATAR AIRWAYS HUNGARIAN GRAND PRIX","link":"https://www.formula1.com/en/results/2023/races/1215/hungary/race-result","year":2023},{"id":34,"name":"MSC CRUISES BELGIAN GRAND PRIX","link":"https://www.formula1.com/en/results/2023/races/1216/belgium/race-result","year":2023},{"id":35,"name":"HEINEKEN DUTCH GRAND PRIX","link":"https://www.formula1.com/en/results/2023/races/1217/netherlands/race-result","year":2023},{"id":36,"name":"PIRELLI GRAN PREMIO D’ITALIA","link":"https://www.formula1.com/en/results/2023/races/1218/italy/race-result","year":2023},{"id":37,"name":"SINGAPORE AIRLINES SINGAPORE GRAND PRIX","link":"https://www.formula1.com/en/results/2023/races/1219/singapore/race-result","year":2023},{"id":38,"name":"LENOVO JAPANESE GRAND PRIX","link":"https://www.formula1.com/en/results/2023/races/1220/japan/race-result","year":2023},{"id":39,"name":"QATAR AIRWAYS QATAR GRAND PRIX","link":"https://www.formula1.com/en/results/2023/races/1221/qatar/race-result","year":2023},{"id":40,"name":"LENOVO UNITED STATES GRAND PRIX","link":"https://www.formula1.com/en/results/2023/races/1222/united-states/race-result","year":2023},{"id":41,"name":"GRAN PREMIO DE LA CIUDAD DE MÉXICO","link":"https://www.formula1.com/en/results/2023/races/1223/mexico/race-result","year":2023},{"id":42,"name":"ROLEX GRANDE PRÊMIO DE SÃO PAULO","link":"https://www.formula1.com/en/results/2023/races/1224/brazil/race-result","year":2023},{"id":43,"name":"HEINEKEN SILVER LAS VEGAS GRAND PRIX","link":"https://www.formula1.com/en/results/2023/races/1225/las-vegas/race-result","year":2023},{"id":44,"name":"ETIHAD AIRWAYS ABU DHABI GRAND PRIX","link":"https://www.formula1.com/en/results/2023/races/1226/abu-dhabi/race-result","year":2023}]}`,
		},
		{
			name:           "Empty response",
			url:            "/getTracks/2020",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"code":404,"message":"error","detail":"not found"}`,
		},
		{
			name:           "Bad param request",
			url:            "/getTracks/alpha",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"code":400,"message":"error","detail":"invalid parameter YEAR"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Config{repo}
			handler := app.routes()

			rr := httptest.NewRecorder()
			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			handler.ServeHTTP(rr, req)

			assertStatusCode(t, rr.Code, tt.expectedStatus)
			assertBody(t, rr.Body.String(), tt.expectedBody)
			assertContentType(t, rr.Header().Get("Content-Type"))
		})
	}
}

func TestGetTrack(t *testing.T) {
	godotenv.Load(".test.env")
	dsn := os.Getenv("DSN")

	db, err := initPostgresDB(dsn)
	if err != nil {
		panic(err)
	}

	repo := data.NewRepository(db)

	tests := []struct {
		name           string
		url            string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Successful request",
			url:            "/getTrack/2023/GULF AIR BAHRAIN GRAND PRIX",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"code":200,"message":"success","detail":"","data":{"id":23,"name":"GULF AIR BAHRAIN GRAND PRIX","link":"https://www.formula1.com/en/results/2023/races/1141/bahrain/race-result","year":2023}}`,
		},
		{
			name:           "Empty response or track not found",
			url:            "/getTrack/2020/invalid",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"code":404,"message":"error","detail":"not found"}`,
		},
		{
			name:           "Bad param request YEAR",
			url:            "/getTrack/alpha/GULF AIR BAHRAIN GRAND PRIX",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"code":400,"message":"error","detail":"invalid parameter YEAR"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Config{repo}
			handler := app.routes()

			rr := httptest.NewRecorder()
			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			handler.ServeHTTP(rr, req)

			assertStatusCode(t, rr.Code, tt.expectedStatus)
			assertBody(t, rr.Body.String(), tt.expectedBody)
			assertContentType(t, rr.Header().Get("Content-Type"))
		})
	}
}

func TestGetResults(t *testing.T) {
	godotenv.Load(".test.env")
	dsn := os.Getenv("DSN")

	db, err := initPostgresDB(dsn)
	if err != nil {
		panic(err)
	}

	repo := data.NewRepository(db)

	tests := []struct {
		name           string
		url            string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Empty response",
			url:            "/getResults/1933",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"code":404,"message":"error","detail":"not found"}`,
		},
		{
			name:           "Bad param request YEAR",
			url:            "/getResults/alpha",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"code":400,"message":"error","detail":"invalid parameter YEAR"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Config{repo}
			handler := app.routes()

			rr := httptest.NewRecorder()
			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			handler.ServeHTTP(rr, req)

			assertStatusCode(t, rr.Code, tt.expectedStatus)
			assertBody(t, rr.Body.String(), tt.expectedBody)
			assertContentType(t, rr.Header().Get("Content-Type"))
		})
	}

	t.Run("Successful request", func(t *testing.T) {
		app := &Config{repo}
		handler := app.routes()

		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/getResults/2023", nil)
		if err != nil {
			t.Fatal(err)
		}

		handler.ServeHTTP(rr, req)

		assertStatusCode(t, rr.Code, http.StatusOK)
		assertContentType(t, rr.Header().Get("Content-Type"))
	})
}

func TestGetResult(t *testing.T) {
	godotenv.Load(".test.env")
	dsn := os.Getenv("DSN")

	db, err := initPostgresDB(dsn)
	if err != nil {
		panic(err)
	}

	repo := data.NewRepository(db)

	tests := []struct {
		name           string
		url            string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Empty response",
			url:            "/getResult/2023/1",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"code":404,"message":"error","detail":"not found"}`,
		},
		{
			name:           "Bad param request year",
			url:            "/getResult/alpha/30",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"code":400,"message":"error","detail":"invalid parameter YEAR"}`,
		},
		{
			name:           "Bad param request track id",
			url:            "/getResult/2023/a",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"code":400,"message":"error","detail":"invalid parameter TRACK_ID"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Config{repo}
			handler := app.routes()

			rr := httptest.NewRecorder()
			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			handler.ServeHTTP(rr, req)

			assertStatusCode(t, rr.Code, tt.expectedStatus)
			assertBody(t, rr.Body.String(), tt.expectedBody)
			assertContentType(t, rr.Header().Get("Content-Type"))
		})
	}

	t.Run("Successful request", func(t *testing.T) {
		app := &Config{repo}
		handler := app.routes()

		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/getResult/2023/30", nil)
		if err != nil {
			t.Fatal(err)
		}

		handler.ServeHTTP(rr, req)

		assertStatusCode(t, rr.Code, http.StatusOK)
		assertContentType(t, rr.Header().Get("Content-Type"))
	})
}
