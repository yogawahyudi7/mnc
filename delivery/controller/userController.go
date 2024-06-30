package controller

import (
	"net/http"

	"github.com/yogawahyudi7/mnc/dto"
	"github.com/yogawahyudi7/mnc/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{userUsecase}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	response, err := c.userUsecase.RegisterUser(req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status": "SUCCESS",
		"result": response,
	})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	response, err := c.userUsecase.Login(req)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status": "SUCCESS",
		"result": response,
	})
}
