package utils

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}
