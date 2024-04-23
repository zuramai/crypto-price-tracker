package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

const LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour

func GenerateToken(email string, secret string) (string, error) {
	claims := JWTClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
