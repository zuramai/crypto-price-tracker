package model

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AuthSuccessResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
