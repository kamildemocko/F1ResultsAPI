package data

import "database/sql"

type Track struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
	Year int    `json:"year"`
}

type Result struct {
	Position      int    `json:"position"`
	DriverNo      int    `json:"driver_no"`
	Driver        string `json:"driver"`
	Car           string `json:"car"`
	Laps          int    `json:"laps"`
	TimeOrRetired string `json:"time_or_retired"`
	Points        int    `json:"points"`
	TrackName     string `json:"track_name"`
	TrackId       int64  `json:"track_id"`
}

func NewRepository(db *sql.DB) Repository {
	return newPostgresRepository(db)
}
