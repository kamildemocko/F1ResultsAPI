package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = 80

type Config struct{}

func main() {
	app := Config{}

	log.Println("starting F1 Results API")

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", webPort),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
