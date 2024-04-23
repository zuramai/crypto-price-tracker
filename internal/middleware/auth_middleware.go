package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zuramai/crypto-price-tracker/internal/services"
)

func NewAuth(authService *services.AuthService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Get token from header
		token := ctx.Get("Authorization")
		if token == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		if len(token) < 7 || token[:7] != "Bearer " {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		token = token[7:]

		// Validate token
		auth, err := authService.ValidateToken(token)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		ctx.Locals("token", token)
		ctx.Locals("auth", auth)

		return ctx.Next()
	}
}
