package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zuramai/crypto-price-tracker/internal/model"
	"github.com/zuramai/crypto-price-tracker/internal/services"
	"github.com/zuramai/crypto-price-tracker/internal/utils"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService,
	}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		return utils.NewApiResponseMessage(ctx, fiber.StatusBadRequest, "Invalid request")
	}

	token, err := c.authService.Login(request.Email, request.Password)
	if err != nil {
		return utils.NewApiResponseMessage(ctx, fiber.StatusUnauthorized, "Invalid credentials")
	}
	response := model.AuthSuccessResponse{
		Email: request.Email,
		Token: *token,
	}
	return utils.NewApiResponse(ctx, fiber.StatusOK, "Login Success", response)
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	request := new(model.LoginUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		return utils.NewApiResponseMessage(ctx, fiber.StatusBadRequest, "Invalid request")
	}

	// Check password confirmation
	if request.Password != request.PasswordConfirmation {
		return utils.NewApiResponseMessage(ctx, fiber.StatusUnprocessableEntity, "Password confirmation doesn't match")

	}

	token, err := c.authService.Register(request.Email, request.Password)
	if err != nil {
		return utils.NewApiResponseMessage(ctx, fiber.StatusUnauthorized, err.Error())
	}
	response := model.AuthSuccessResponse{
		Email: request.Email,
		Token: token,
	}
	return utils.NewApiResponse(ctx, fiber.StatusOK, "Register Success", response)
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {

	ctx.Locals("auth")
	// invalidate token
	return nil
}
