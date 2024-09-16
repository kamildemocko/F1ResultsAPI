package data

import "database/sql"

type Repository interface {
	GetTracks(int) (*[]Track, error)
	GetTrack(int, string) (*Track, error)
	GetResults(int) (*[]Result, error)
	GetResult(int, int64) (*[]Result, error)
}

type postgresRepository struct {
	DB *sql.DB
}

func newPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{DB: db}
}
