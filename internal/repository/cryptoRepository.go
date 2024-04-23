package repository

import (
	"database/sql"
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

func (repo *CryptoRepository) FindAll() []*model.Crypto {
	var result []*model.Crypto
	return result
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
