package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yogawahyudi7/mnc/delivery/controller"
)

func RegisterAuthRoutes(app fiber.Router, authController *controller.AuthController) {
	app.Post("/register", authController.Register)
	app.Post("/login", authController.Login)
}
