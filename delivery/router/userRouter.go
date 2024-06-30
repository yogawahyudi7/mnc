package router

import (
	"github.com/yogawahyudi7/mnc/delivery/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, userController *controller.UserController) {
	user := app.Group("/user")
	user.Post("/register", userController.Register)
	user.Post("/login", userController.Login)
}
