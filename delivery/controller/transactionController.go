package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yogawahyudi7/mnc/dto"
	"github.com/yogawahyudi7/mnc/pkg/constant"
	"github.com/yogawahyudi7/mnc/usecase"
)

type transactionController struct {
	transactionUsecase usecase.TransactionUsecase
}

func NewTransactionController(transactionUsecase usecase.TransactionUsecase) *transactionController {
	return &transactionController{transactionUsecase}
}

func (c *transactionController) TopUp(ctx *fiber.Ctx) error {

	uuid := ctx.Locals("uuid").(string)

	var req dto.TopUpRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": constant.InvalidRequestBody,
		})
	}

	response, err := c.transactionUsecase.TopUp(req, uuid)
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
