package main

import (
	"github.com/yogawahyudi7/mnc/config"
	"github.com/yogawahyudi7/mnc/delivery/controller"
	"github.com/yogawahyudi7/mnc/delivery/router"
	"github.com/yogawahyudi7/mnc/pkg/database"
	"github.com/yogawahyudi7/mnc/repository"
	"github.com/yogawahyudi7/mnc/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// setup
	setup := &config.Server{}
	setup.Load()

	// database
	db := database.NewDatabase(setup)

	// repository
	tokenRepository := repository.NewTokenRepository(db)
	userRepository := repository.NewUserRepository(db)

	// usecase
	authUsecase := usecase.NewAuthUsecase(setup, userRepository, tokenRepository)
	userUsecase := usecase.NewUserUsecase(setup, userRepository, tokenRepository)

	// controller
	authController := controller.NewAuthController(authUsecase)
	userController := controller.NewUserController(userUsecase)

	// router
	app := fiber.New()
	authGroup := app.Group("/auth")
	router.RegisterAuthRoutes(authGroup, authController)

	userGroup := app.Group("/user")
	router.RegisterUserRoutes(userGroup, setup, userController)

	// run app
	app.Listen(setup.AppPort)
}
