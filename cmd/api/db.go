package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	maxOpenDbConn = 25
	maxIdleDBConn = 25
	maxDBLifetime = 5 * time.Minute
)

func initPostgresDB() (*sql.DB, error) {
	dsn := os.Getenv("DSN")

	log.Println("connecting to DB")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDBConn)
	db.SetConnMaxLifetime(maxDBLifetime)

	return db, nil
}
