package utils

import "github.com/gofiber/fiber/v2"

type ApiResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}
type ApiResponseMessage struct {
	Message string `json:"message"`
}

func NewApiResponseMessage(ctx *fiber.Ctx, code int, message string) error {
	return ctx.Status(code).JSON(&ApiResponseMessage{
		Message: message,
	})
}
func NewApiResponse[T any](ctx *fiber.Ctx, code int, message string, data T) error {
	return ctx.Status(code).JSON(&ApiResponse[T]{
		Message: message,
		Data:    data,
	})
}
