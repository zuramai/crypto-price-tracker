package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zuramai/crypto-price-tracker/internal/services"
)

type TrackerController struct {
	trackerService *services.TrackerService
}

func NewTrackerController(trackerService *services.TrackerService) *TrackerController {
	return &TrackerController{trackerService}
}

func (c *TrackerController) GetTrackers(ctx *fiber.Ctx) error {
	return nil
}
func (c *TrackerController) CreateTracker(ctx *fiber.Ctx) error {
	return nil
}
func (c *TrackerController) DeleteTracker(ctx *fiber.Ctx) error {
	return nil
}
