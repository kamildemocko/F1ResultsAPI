package data

import (
	"context"
	"time"
)

func (pg *postgresRepository) GetTracks(year int) (Tracks, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, name, link, year
	FROM f1scrap.tracks
	WHERE year = $1;`

	rows, err := pg.DB.QueryContext(ctx, query, year)
	if err != nil {
		return nil, err
	}

	var tracks []Track

	for rows.Next() {
		var track Track
		rows.Scan(&track.ID, &track.Name, &track.Link, &track.Year)

		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (pg *postgresRepository) GetTrack(year int, trackName string) (Track, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, name, link, year
	FROM f1scrap.tracks
	WHERE year = $1 AND name;`

	rows := pg.DB.QueryRowContext(ctx, query, year)

	var track Track
	err := rows.Scan(&track.ID, &track.Name, &track.Link, &track.Year)
	if err != nil {
		return Track{}, err
	}

	return track, nil
}
