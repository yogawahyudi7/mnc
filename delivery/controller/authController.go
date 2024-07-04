package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yogawahyudi7/mnc/dto"
	"github.com/yogawahyudi7/mnc/pkg/constant"
	"github.com/yogawahyudi7/mnc/usecase"
)

type AuthController struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthController(userUsecase usecase.AuthUsecase) *AuthController {
	return &AuthController{userUsecase}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": constant.InvalidRequestBody,
		})
	}

	response, err := c.authUsecase.RegisterUser(req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status": constant.UpperCaseSuccess,
		"result": response,
	})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": constant.InvalidRequestBody,
		})
	}

	response, err := c.authUsecase.Login(req)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status": constant.UpperCaseSuccess,
		"result": response,
	})
}
