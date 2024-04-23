package utils

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

func GenerateToken(email string, secret string) (string, error) {
	claims := JWTClaims{
		Email: email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
