package main

import (
	"F1ResultsApi/data"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

const webPort = 80

type Config struct {
	Repository data.Repository
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	app := Config{}

	// create db
	db, err := initPostgresDB()
	if err != nil {
		panic(err)
	}

	app.Repository = data.NewRepository(db)

	log.Println("starting F1 Results API")

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", webPort),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
