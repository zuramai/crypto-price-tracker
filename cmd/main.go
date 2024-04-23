package main

import (
	"fmt"

	"github.com/zuramai/crypto-price-tracker/internal/config"
	"github.com/zuramai/crypto-price-tracker/internal/controller"
	"github.com/zuramai/crypto-price-tracker/internal/middleware"
	"github.com/zuramai/crypto-price-tracker/internal/repository"
	"github.com/zuramai/crypto-price-tracker/internal/router"
	"github.com/zuramai/crypto-price-tracker/internal/services"
)

func main() {
	viper := config.NewViper()
	app := config.NewFiber(viper)
	logger := config.NewLogger(viper)
	db := config.NewDatabase(viper, logger)

	userRepo := repository.NewUserRepository(db, logger)
	trackerRepo := repository.NewTrackerRepository(db)
	// cryptoRepo := repository.NewCryptoRepository(db)

	authService := services.NewAuthService(userRepo, logger, viper)
	trackerService := services.NewTrackerService(trackerRepo, userRepo)

	authController := controller.NewAuthController(authService)
	trackerController := controller.NewTrackerController(trackerService)
	homeController := controller.NewHomeController()

	authMiddleware := middleware.NewAuth(authService)

	router := router.RouteConfig{
		App:               app,
		AuthController:    authController,
		TrackerController: trackerController,
		AuthMiddleware:    authMiddleware,
		HomeController:    homeController,
	}

	router.RegisterRoutes()

	port := viper.GetString("app.port")
	logger.Infof("App running at port %v", port)
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		logger.Fatalf("Failed to start the server: %v", err)
	}
}
