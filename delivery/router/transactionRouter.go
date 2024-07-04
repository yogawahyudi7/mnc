package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yogawahyudi7/mnc/config"
	"github.com/yogawahyudi7/mnc/delivery/controller"
	"github.com/yogawahyudi7/mnc/delivery/middleware"
)

func RegisterTransactionRoutes(router fiber.Router, config *config.Server, controller *controller.TransactionController) {
	router.Post("/topup", middleware.Authentication(config), controller.TopUp)
	router.Post("/payment", middleware.Authentication(config), controller.Payment)
	router.Post("/transfer", middleware.Authentication(config), controller.Transfer)
}
