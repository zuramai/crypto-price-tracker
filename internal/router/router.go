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

	// Guest routes
	r.App.Post("/auth/login", r.AuthController.Login)
	r.App.Post("/auth/register", r.AuthController.Register)

	// Auth Routes
	r.App.Post("/auth/logout", r.AuthMiddleware, r.AuthController.Logout)
	r.App.Get("/crypto", r.AuthMiddleware, r.CryptoController.GetCryptos)
	r.App.Get("/trackers", r.AuthMiddleware, r.TrackerController.GetTrackers)
	r.App.Post("/trackers", r.AuthMiddleware, r.TrackerController.CreateTracker)
	r.App.Delete("/trackers/:id", r.AuthMiddleware, r.TrackerController.DeleteTracker)
}
