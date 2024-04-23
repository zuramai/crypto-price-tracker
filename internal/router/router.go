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
	HomeController    *controller.HomeController
	AuthMiddleware    fiber.Handler
}

func (r *RouteConfig) RegisterRoutes() {
	r.App.Get("/", r.HomeController.Home)

	api := r.App.Group("/api")

	// Guest routes
	api.Post("/auth/login", r.AuthController.Login)
	api.Post("/auth/register", r.AuthController.Register)

	// Auth Routes
	api.Post("/auth/logout", r.AuthMiddleware, r.AuthController.Logout)
	api.Get("/crypto", r.AuthMiddleware, r.CryptoController.GetCryptos)
	api.Get("/trackers", r.AuthMiddleware, r.TrackerController.GetTrackers)
	api.Post("/trackers", r.AuthMiddleware, r.TrackerController.CreateTracker)
	api.Delete("/trackers/:id", r.AuthMiddleware, r.TrackerController.DeleteTracker)
}
