package services

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/zuramai/crypto-price-tracker/internal/model"
	"github.com/zuramai/crypto-price-tracker/internal/repository"
	"go.uber.org/zap"
)

type CryptoService struct {
	cryptoRepo *repository.CryptoRepository
	logger     *zap.SugaredLogger
}

func NewCryptoService(cryptoRepo *repository.CryptoRepository, logger *zap.SugaredLogger) *CryptoService {
	return &CryptoService{cryptoRepo, logger}
}

func (s *CryptoService) fetchCryptos() []model.Crypto {
	resp, err := http.Get("https://api.coincap.io/v2/assets?limit=2000")
	if err != nil {
		s.logger.Errorf("error fetching crypto data: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Errorf("error reading response body: %v", err)
	}

	var response struct {
		Data      []model.Crypto `json:"data"`
		Timestamp int            `json:"timestamp"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		s.logger.Errorf("error unmarshalling response body: %v", err)
	}

	return response.Data
}

func (s *CryptoService) FetchCryptoContinuous() {
	for {
		time.Sleep(5 * time.Second)
		s.logger.Debugf("Fetching crypto data...")
		cryptos := s.fetchCryptos()
		err := s.cryptoRepo.UpdateCryptos(cryptos)
		if err != nil {
			s.logger.Errorf("error updating cryptos: %v", err)
		}

	}
}
