package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yogawahyudi7/mnc/dto"
	"github.com/yogawahyudi7/mnc/pkg/constant"
	"github.com/yogawahyudi7/mnc/usecase"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{userUsecase}
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	claims := ctx.Locals(constant.UserContext).(jwt.MapClaims)

	id := claims["id"].(string)
	response, err := c.userUsecase.GetUser(id)
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

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	claims := ctx.Locals(constant.UserContext).(jwt.MapClaims)

	userId := claims["id"].(string)
	var req dto.UpdateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": constant.InvalidRequestBody,
		})
	}

	response, err := c.userUsecase.UpdateUser(req, userId)
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
