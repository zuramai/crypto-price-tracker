package services

import (
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"github.com/zuramai/crypto-price-tracker/internal/model"
	"github.com/zuramai/crypto-price-tracker/internal/repository"
	"github.com/zuramai/crypto-price-tracker/internal/utils"
	"go.uber.org/zap"
)

type AuthService struct {
	userRepo *repository.UserRepository
	logger   *zap.SugaredLogger
	viper    *viper.Viper
}

func NewAuthService(userRepo *repository.UserRepository, logger *zap.SugaredLogger, viper *viper.Viper) *AuthService {
	return &AuthService{userRepo, logger, viper}
}

func (s *AuthService) ValidateToken(token string) (*model.User, error) {
	var user *model.User

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.viper.GetString("app.jwt_secret")), nil
	})

	if err != nil {
		return nil, err
	}

	claims := t.Claims.(*utils.JWTClaims)

	user, err = s.userRepo.FindUserByEmail(claims.Email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
