package repository

import (
	"database/sql"

	"github.com/zuramai/crypto-price-tracker/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	return &model.User{}, nil
}
