package model

type Tracker struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	CryptoID string `json:"crypto_id"`
	User     User   `json:"user"`
	Crypto   Crypto `json:"crypto"`
}

type CreateTrackerRequest struct {
	CryptoID string `json:"crypto_id"`
}
