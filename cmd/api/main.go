package main

import (
	"F1ResultsApi/data"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const webPort = 80

var dsn string

type Config struct {
	Repository data.Repository
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env not found, using default envirnmental variables")
	}

	dsn = os.Getenv("DSN")
	cwd, err := os.Getwd()
	fmt.Println(cwd)
}

func main() {
	app := Config{}

	// create db
	db, err := initPostgresDB(dsn)
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
