package repository

import (
	"database/sql"
	"errors"

	"github.com/zuramai/crypto-price-tracker/internal/model"
	"go.uber.org/zap"
)

type UserRepository struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

func NewUserRepository(db *sql.DB, logger *zap.SugaredLogger) *UserRepository {
	return &UserRepository{db, logger}
}

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrInsertUser        = errors.New("failed to insert user")
	ErrQueryUser         = errors.New("failed to query user")
)

func (repo *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	rows := repo.db.QueryRow("SELECT * FROM users WHERE email = $1 LIMIT 1", email)

	err := rows.Scan(&user.ID, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, ErrQueryUser
	}

	return &user, nil
}
func (repo *UserRepository) CreateUser(email string, password string) error {
	user, err := repo.FindUserByEmail(email)
	repo.logger.Debugf("user: %v", user)
	if err == nil {
		return ErrUserAlreadyExists
	}

	_, err = repo.db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, password)

	if err != nil {
		repo.logger.Errorf("error insert user: %v", err)
		return ErrInsertUser
	}
	return nil
}
