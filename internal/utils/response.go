package utils

type ApiResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}
