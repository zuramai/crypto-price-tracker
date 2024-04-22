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

		// Validate token
		err := authService.ValidateToken(token)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		ctx.UserContext()

		ctx.Locals("auth", auth)

		return ctx.Next()
	}
}
