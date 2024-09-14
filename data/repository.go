package data

import "database/sql"

type Repository interface {
	GetTracks(int) (Tracks, error)
	GetTrack(int, string) (Track, error)
}

type postgresRepository struct {
	DB *sql.DB
}

func newPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{DB: db}
}
