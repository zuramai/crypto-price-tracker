package repository

import (
	"database/sql"

	"github.com/zuramai/crypto-price-tracker/internal/model"
)

type CryptoRepository struct {
	db *sql.DB
}

func NewCryptoRepository(db *sql.DB) *CryptoRepository {
	return &CryptoRepository{db}
}

func (repo *CryptoRepository) FindAll() []*model.Crypto {
	var result []*model.Crypto
	return result
}
