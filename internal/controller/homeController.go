package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zuramai/crypto-price-tracker/internal/utils"
)

type HomeController struct {
}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (c *HomeController) Home(ctx *fiber.Ctx) error {
	return ctx.JSON(utils.ApiResponse[fiber.Map]{
		Message: "Welcome to Crypto Price Tracker!",
		Data:    fiber.Map{},
	})
}
