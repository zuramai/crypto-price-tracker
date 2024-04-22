package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zuramai/crypto-price-tracker/internal/services"
)

type AuthController struct {
	userService *services.AuthService
}

func NewAuthController(userService *services.AuthService) *AuthController {
	return &AuthController{
		userService,
	}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	return nil
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	return nil
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	return nil
}
