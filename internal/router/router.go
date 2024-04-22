package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zuramai/crypto-price-tracker/internal/controller"
)

type RouteConfig struct {
	App               *fiber.App
	AuthController    *controller.AuthController
	CryptoController  *controller.CryptoController
	TrackerController *controller.TrackerController
	AuthMiddleware    fiber.Handler
}

func (r *RouteConfig) RegisterRoutes() {
	// Guest routes
	r.App.Post("/auth/login", r.AuthController.Login)
	r.App.Post("/auth/register", r.AuthController.Register)

	// Auth Routes
	r.App.Use(r.AuthMiddleware)
	r.App.Post("/auth/logout", r.AuthController.Logout)
	r.App.Get("/crypto", r.CryptoController.GetCryptos)
	r.App.Get("/trackers", r.TrackerController.GetTrackers)
	r.App.Post("/trackers", r.TrackerController.CreateTracker)
	r.App.Delete("/trackers/:id", r.TrackerController.DeleteTracker)
}
