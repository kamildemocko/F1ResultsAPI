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

func (pg *postgresRepository) GetResults(year int) (Results, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT "position", driver_no, driver, car, laps, time_or_retired, points, name as track_name, track_id
	FROM f1scrap.tracks
	JOIN f1scrap.results ON tracks.id = results.track_id
	WHERE year = $1;`

	rows, err := pg.DB.QueryContext(ctx, query, year)
	if err != nil {
		return nil, err
	}

	var results []Result

	for rows.Next() {
		var rs Result
		rows.Scan(
			&rs.Position, &rs.DriverNo, &rs.Driver, &rs.Car, &rs.Laps,
			&rs.TimeOrRetired, &rs.Points, &rs.TrackName, &rs.TrackId,
		)
		results = append(results, rs)
	}

	return results, nil
}
