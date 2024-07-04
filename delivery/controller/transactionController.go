package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yogawahyudi7/mnc/dto"
	"github.com/yogawahyudi7/mnc/pkg/constant"
	"github.com/yogawahyudi7/mnc/usecase"
)

type TransactionController struct {
	transactionUsecase usecase.TransactionUsecase
}

func NewTransactionController(transactionUsecase usecase.TransactionUsecase) *TransactionController {
	return &TransactionController{transactionUsecase}
}

func (c *TransactionController) TopUp(ctx *fiber.Ctx) error {

	claims := ctx.Locals(constant.UserContext).(jwt.MapClaims)
	id := claims["id"].(string)

	var req dto.TopUpRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": constant.InvalidRequestBody,
		})
	}

	response, err := c.transactionUsecase.TopUp(req, id)
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

func (c *TransactionController) Payment(ctx *fiber.Ctx) error {

	claims := ctx.Locals(constant.UserContext).(jwt.MapClaims)
	id := claims["id"].(string)

	var req dto.PaymentRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": constant.InvalidRequestBody,
		})
	}

	response, err := c.transactionUsecase.Payment(req, id)
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

func (c *TransactionController) Transfer(ctx *fiber.Ctx) error {

	claims := ctx.Locals(constant.UserContext).(jwt.MapClaims)
	id := claims["id"].(string)

	var req dto.TransferRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": constant.InvalidRequestBody,
		})
	}

	response, err := c.transactionUsecase.Transfer(req, id)
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
