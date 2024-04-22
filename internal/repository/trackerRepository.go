package repository

import "database/sql"

type TrackerRepository struct {
	db *sql.DB
}

func NewTrackerRepository(db *sql.DB) *TrackerRepository {
	return &TrackerRepository{db}
}
