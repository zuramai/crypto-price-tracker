package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type TrackerRepository struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

func NewTrackerRepository(db *sql.DB, logger *zap.SugaredLogger) *TrackerRepository {
	return &TrackerRepository{db, logger}
}
