package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zuramai/crypto-price-tracker/internal/model"
	"github.com/zuramai/crypto-price-tracker/internal/services"
	"github.com/zuramai/crypto-price-tracker/internal/utils"
)

type TrackerController struct {
	trackerService *services.TrackerService
}

func NewTrackerController(trackerService *services.TrackerService) *TrackerController {
	return &TrackerController{trackerService}
}

func (c *TrackerController) GetTrackers(ctx *fiber.Ctx) error {
	user := ctx.Locals("auth").(*model.User)
	trackers, err := c.trackerService.GetTrackersByUserID(user.ID)
	if err != nil {
		return utils.NewApiResponseMessage(ctx, fiber.StatusInternalServerError, "Failed to get trackers. Internal server error.")
	}
	return utils.NewApiResponse(ctx, fiber.StatusOK, "Success get trackers", trackers)
}
func (c *TrackerController) CreateTracker(ctx *fiber.Ctx) error {
	user := ctx.Locals("auth").(*model.User)
	var request model.CreateTrackerRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return utils.NewApiResponseMessage(ctx, fiber.StatusBadRequest, "Invalid request")
	}

	err = c.trackerService.CreateTracker(user.ID, request.CryptoID)
	if err != nil {
		return utils.NewApiResponseMessage(ctx, fiber.StatusBadRequest, err.Error())
	}
	return utils.NewApiResponseMessage(ctx, fiber.StatusCreated, "Tracker created successfully")
}
func (c *TrackerController) DeleteTracker(ctx *fiber.Ctx) error {
	trackerID, err := ctx.ParamsInt("id")
	if err != nil {
		return utils.NewApiResponseMessage(ctx, fiber.StatusBadRequest, "Tracker id should be number")
	}
	err = c.trackerService.DeleteTracker(trackerID)
	if err != nil {
		return utils.NewApiResponseMessage(ctx, fiber.StatusBadRequest, err.Error())
	}
	return utils.NewApiResponseMessage(ctx, fiber.StatusOK, "Tracker deleted successfully")
}
