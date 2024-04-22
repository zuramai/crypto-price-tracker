package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zuramai/crypto-price-tracker/internal/repository"
)

type AuthController struct {
	userRepo *repository.UserRepository
}

func NewAuthController(userRepo *repository.UserRepository) *AuthController {
	return &AuthController{
		userRepo,
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
