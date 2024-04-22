package services

import (
	"github.com/zuramai/crypto-price-tracker/internal/repository"
	"go.uber.org/zap"
)

type AuthService struct {
	userRepo *repository.UserRepository
	logger   *zap.SugaredLogger
}

func NewAuthService(userRepo *repository.UserRepository, logger *zap.SugaredLogger) *AuthService {
	return &AuthService{userRepo, logger}
}

func (s *AuthService) ValidateToken(token string) error {
	return nil
}
