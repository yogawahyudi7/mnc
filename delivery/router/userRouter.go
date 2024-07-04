package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yogawahyudi7/mnc/config"
	"github.com/yogawahyudi7/mnc/delivery/controller"
	"github.com/yogawahyudi7/mnc/delivery/middleware"
)

func RegisterUserRoutes(app fiber.Router, config *config.Server, userController *controller.UserController) {
	app.Get("/profile", middleware.Authentication(config), userController.GetUser)
	app.Put("/update", middleware.Authentication(config), userController.UpdateUser)
}
