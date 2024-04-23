package services

import (
	"errors"

	"github.com/zuramai/crypto-price-tracker/internal/model"
	"github.com/zuramai/crypto-price-tracker/internal/repository"
)

type TrackerService struct {
	userRepo    *repository.UserRepository
	trackerRepo *repository.TrackerRepository
}

func NewTrackerService(trackerRepo *repository.TrackerRepository, userRepo *repository.UserRepository) *TrackerService {
	return &TrackerService{userRepo, trackerRepo}
}

var (
	ErrTrackerAlreadyExists = errors.New("tracker already exists")
	ErrInsertTracker        = errors.New("failed to create tracker")
)

func (s *TrackerService) GetTrackersByUserID(userID int) ([]model.Tracker, error) {
	trackers, err := s.trackerRepo.GetTrackersByUserID(userID)
	if err != nil {
		return nil, err
	}
	return trackers, nil
}
func (s *TrackerService) CreateTracker(userID int, cryptoID string) error {
	tracker, err := s.trackerRepo.FindTracker(userID, cryptoID)
	if err != nil && err != repository.ErrTrackerNotFound {
		return err
	}
	if tracker != nil {
		return ErrTrackerAlreadyExists
	}

	err = s.trackerRepo.InsertTracker(userID, cryptoID)
	if err != nil {
		return ErrInsertTracker
	}

	return nil
}
