package repository

import (
	"database/sql"
	"errors"

	"github.com/zuramai/crypto-price-tracker/internal/model"
	"go.uber.org/zap"
)

type TrackerRepository struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

func NewTrackerRepository(db *sql.DB, logger *zap.SugaredLogger) *TrackerRepository {
	return &TrackerRepository{db, logger}
}

var (
	ErrQueryTracker    = errors.New("failed to query tracker")
	ErrTrackerNotFound = errors.New("tracker not found")
	ErrInsertTracker   = errors.New("failed to insert tracker")
)

func (repo *TrackerRepository) GetTrackersByUserID(userID int) ([]model.Tracker, error) {
	var trackers []model.Tracker
	rows, err := repo.db.Query("SELECT * FROM trackers WHERE user_id = $1", userID)
	if err != nil {
		repo.logger.Errorf("error query trackers: %v", err)
		return nil, ErrQueryTracker
	}
	rows.Scan(&trackers)
	if err != nil {
		repo.logger.Errorf("error scan trackers: %v", err)
		return nil, ErrQueryTracker
	}
	return trackers, nil
}

func (repo *TrackerRepository) FindTracker(userID int, cryptoID string) (*model.Tracker, error) {
	var tracker model.Tracker
	rows := repo.db.QueryRow("SELECT * FROM trackers WHERE user_id = $1", userID)
	if rows.Err() != nil {
		if rows.Err() == sql.ErrNoRows {
			return nil, ErrTrackerNotFound
		}
		repo.logger.Errorf("error query  trackers: %v", rows.Err())
		return nil, rows.Err()
	}
	err := rows.Scan(&tracker.ID, &tracker.UserID, &tracker.CryptoID)

	if err != nil {
		repo.logger.Errorf("error scan trackers: %v", err)
		if err == sql.ErrNoRows {
			return nil, ErrTrackerNotFound
		}
		return nil, ErrQueryTracker
	}

	return &tracker, nil
}

func (repo *TrackerRepository) InsertTracker(userID int, cryptoID string) error {
	_, err := repo.db.Exec("INSERT INTO trackers (user_id, crypto_id) VALUES ($1, $2)", userID, cryptoID)
	if err != nil {
		repo.logger.Errorf("error insert trackers: %v", err)
		return ErrInsertTracker
	}

	return nil
}
