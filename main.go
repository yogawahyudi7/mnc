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
	setup := &config.Server{}
	setup.Load()

	db := database.NewDatabase(setup)

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	app := fiber.New()
	router.SetupUserRoutes(app, userController)

	app.Listen(setup.AppPort)
}
