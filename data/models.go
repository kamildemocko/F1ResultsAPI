package data

import "database/sql"

type Tracks []Track

type Track struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
	Year int    `json:"year"`
}

type Results []Result

type Result struct{}

func NewRepository(db *sql.DB) Repository {
	return newPostgresRepository(db)
}
