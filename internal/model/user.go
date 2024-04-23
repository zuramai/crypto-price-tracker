package model

import "database/sql"

type User struct {
	ID       int            `json:"id"`
	Email    string         `json:"email"`
	Password string         `json:"-"`
	Token    sql.NullString `json:"-"`
}

type LoginUserRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}
type AuthSuccessResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
