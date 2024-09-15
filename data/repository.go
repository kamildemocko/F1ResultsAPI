package data

import "database/sql"

type Repository interface {
	GetTracks(int) (Tracks, error)
	GetTrack(int, string) (Track, error)
	GetResults(int) (Results, error)
	// GetResult(int, int64) (Results, error)
}

type postgresRepository struct {
	DB *sql.DB
}

func newPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{DB: db}
}
