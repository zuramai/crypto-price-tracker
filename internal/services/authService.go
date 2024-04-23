package services

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"github.com/zuramai/crypto-price-tracker/internal/model"
	"github.com/zuramai/crypto-price-tracker/internal/repository"
	"github.com/zuramai/crypto-price-tracker/internal/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
	logger   *zap.SugaredLogger
	viper    *viper.Viper
}

func NewAuthService(userRepo *repository.UserRepository, logger *zap.SugaredLogger, viper *viper.Viper) *AuthService {
	return &AuthService{userRepo, logger, viper}
}

var (
	ErrInvalidToken = errors.New("invalid token")
)

func (s *AuthService) ValidateToken(token string) (*model.User, error) {
	var user *model.User

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.viper.GetString("app.jwt_secret")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)

	if !ok || !t.Valid {
		return nil, ErrInvalidToken
	}

	user, err = s.userRepo.FindUserByEmail(claims["email"].(string))

	if user.Token.String != token {
		return nil, ErrInvalidToken
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(email string, password string) (*string, error) {
	user, err := s.userRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	// Generate token
	token, err := utils.GenerateToken(email, s.viper.GetString("app.jwt_secret"))
	if err != nil {
		return nil, err
	}

	// Update token in database
	err = s.userRepo.UpdateToken(email, token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *AuthService) Register(email string, password string) (*string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	err = s.userRepo.CreateUser(email, string(hashedPassword[:]))
	if err != nil {
		return nil, err
	}

	// Generate token
	token, err := utils.GenerateToken(email, s.viper.GetString("app.jwt_secret"))
	if err != nil {
		return nil, err
	}

	// Update token in database
	err = s.userRepo.UpdateToken(email, token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *AuthService) Logout(token string) error {
	// Invalidate token
	err := s.userRepo.RemoveToken(token)
	if err != nil {
		return err
	}
	return nil
}
