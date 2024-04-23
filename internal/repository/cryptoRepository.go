package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/zuramai/crypto-price-tracker/internal/model"
	"go.uber.org/zap"
)

type CryptoRepository struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

func NewCryptoRepository(db *sql.DB, logger *zap.SugaredLogger) *CryptoRepository {
	return &CryptoRepository{db, logger}
}

var (
	ErrQueryCrypto    = errors.New("error query crypto")
	ErrCryptoNotFound = errors.New("crypto not found")
)

func (repo *CryptoRepository) FindAll() []*model.Crypto {
	var result []*model.Crypto
	return result
}

func (repo *CryptoRepository) FindCryptoByID(id string) (*model.Crypto, error) {
	var result model.Crypto
	rows := repo.db.QueryRow("SELECT * FROM cryptos WHERE id = $1", id)
	if rows.Err() != nil {
		if rows.Err() == sql.ErrNoRows {
			return nil, ErrCryptoNotFound
		}
		repo.logger.Errorf("error query crypto: %v", rows.Err())
		return nil, rows.Err()
	}
	err := rows.Scan(&result.ID, &result.Rank, &result.Symbol, &result.Name, &result.Supply, &result.MaxSupply, &result.MarketCapUSD, &result.VolumeUsd24Hr, &result.PriceUSD, &result.ChangePercent24Hr, &result.Vwap24Hr)

	if err != nil {
		repo.logger.Errorf("error scan crypto: %v", err)
		if err == sql.ErrNoRows {
			return nil, ErrCryptoNotFound
		}
		return nil, ErrQueryCrypto
	}
	return &result, nil
}
func (repo *CryptoRepository) UpdateCryptos(cryptos []model.Crypto) error {
	var query string
	for _, el := range cryptos {
		query += fmt.Sprintf(
			`INSERT OR REPLACE INTO cryptos (
				id,
				rank,
				symbol,
				name,
				supply,
				maxSupply,
				marketCapUsd,
				volumeUsd24Hr,
				priceUsd,
				changePercent24Hr,
				vwap24Hr
			) VALUES ("%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s");`,
			el.ID, el.Rank, el.Symbol, el.Name, el.Supply, el.MaxSupply, el.MarketCapUSD, el.VolumeUsd24Hr, el.PriceUSD, el.ChangePercent24Hr, el.Vwap24Hr)
	}
	_, err := repo.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
