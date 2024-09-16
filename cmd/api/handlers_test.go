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
			url:            "/getTracks/2024",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"code":200,"message":"success","detail":"","data":[{"id":1,"name":"FORMULA 1 GULF AIR BAHRAIN GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1229/bahrain/race-result","year":2024},{"id":2,"name":"FORMULA 1 STC SAUDI ARABIAN GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1230/saudi-arabia/race-result","year":2024},{"id":3,"name":"FORMULA 1 ROLEX AUSTRALIAN GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1231/australia/race-result","year":2024},{"id":4,"name":"FORMULA 1 MSC CRUISES JAPANESE GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1232/japan/race-result","year":2024},{"id":5,"name":"FORMULA 1 LENOVO CHINESE GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1233/china/race-result","year":2024},{"id":6,"name":"FORMULA 1 CRYPTO.COM MIAMI GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1234/miami/race-result","year":2024},{"id":7,"name":"FORMULA 1 MSC CRUISES GRAN PREMIO DEL MADE IN ITALY E DELL'EMILIA-ROMAGNA 2024","link":"https://www.formula1.com/en/results/2024/races/1235/emilia-romagna/race-result","year":2024},{"id":8,"name":"FORMULA 1 GRAND PRIX DE MONACO 2024","link":"https://www.formula1.com/en/results/2024/races/1236/monaco/race-result","year":2024},{"id":9,"name":"FORMULA 1 AWS GRAND PRIX DU CANADA 2024","link":"https://www.formula1.com/en/results/2024/races/1237/canada/race-result","year":2024},{"id":10,"name":"FORMULA 1 ARAMCO GRAN PREMIO DE ESPAÑA 2024","link":"https://www.formula1.com/en/results/2024/races/1238/spain/race-result","year":2024},{"id":11,"name":"FORMULA 1 QATAR AIRWAYS AUSTRIAN GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1239/austria/race-result","year":2024},{"id":12,"name":"FORMULA 1 QATAR AIRWAYS BRITISH GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1240/great-britain/race-result","year":2024},{"id":13,"name":"FORMULA 1 HUNGARIAN GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1241/hungary/race-result","year":2024},{"id":14,"name":"FORMULA 1 ROLEX BELGIAN GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1242/belgium/race-result","year":2024},{"id":15,"name":"FORMULA 1 HEINEKEN DUTCH GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1243/netherlands/race-result","year":2024},{"id":16,"name":"FORMULA 1 PIRELLI GRAN PREMIO D’ITALIA 2024","link":"https://www.formula1.com/en/results/2024/races/1244/italy/race-result","year":2024},{"id":17,"name":"FORMULA 1 QATAR AIRWAYS AZERBAIJAN GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1245/azerbaijan/race-result","year":2024},{"id":18,"name":"FORMULA 1 SINGAPORE AIRLINES SINGAPORE GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1246/singapore/race-result","year":2024},{"id":19,"name":"FORMULA 1 PIRELLI UNITED STATES GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1247/united-states/race-result","year":2024},{"id":20,"name":"FORMULA 1 GRAN PREMIO DE LA CIUDAD DE MÉXICO 2024","link":"https://www.formula1.com/en/results/2024/races/1248/mexico/race-result","year":2024},{"id":21,"name":"FORMULA 1 LENOVO GRANDE PRÊMIO DE SÃO PAULO 2024","link":"https://www.formula1.com/en/results/2024/races/1249/brazil/race-result","year":2024},{"id":22,"name":"FORMULA 1 HEINEKEN SILVER LAS VEGAS GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1250/las-vegas/race-result","year":2024},{"id":23,"name":"FORMULA 1 QATAR AIRWAYS QATAR GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1251/qatar/race-result","year":2024},{"id":24,"name":"FORMULA 1 ETIHAD AIRWAYS ABU DHABI GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1252/abu-dhabi/race-result","year":2024}]}`,
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
			url:            "/getTrack/2024/FORMULA 1 GULF AIR BAHRAIN GRAND PRIX 2024",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"code":200,"message":"success","detail":"","data":{"id":1,"name":"FORMULA 1 GULF AIR BAHRAIN GRAND PRIX 2024","link":"https://www.formula1.com/en/results/2024/races/1229/bahrain/race-result","year":2024}}`,
		},
		{
			name:           "Empty response or track not found",
			url:            "/getTrack/2020/invalid",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"code":404,"message":"error","detail":"not found"}`,
		},
		{
			name:           "Bad param request YEAR",
			url:            "/getTrack/alpha/FORMULA 1 GULF AIR BAHRAIN GRAND PRIX 2024",
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
