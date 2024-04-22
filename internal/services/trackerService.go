package services

import (
	"github.com/zuramai/crypto-price-tracker/internal/repository"
)

type TrackerService struct {
	userRepo    *repository.UserRepository
	trackerRepo *repository.TrackerRepository
}

func NewTrackerService(trackerRepo *repository.TrackerRepository, userRepo *repository.UserRepository) *TrackerService {
	return &TrackerService{userRepo, trackerRepo}
}
