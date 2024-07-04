package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yogawahyudi7/mnc/config"
	"github.com/yogawahyudi7/mnc/pkg/constant"
	"github.com/yogawahyudi7/mnc/pkg/jwt"
)

func Authentication(config *config.Server) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get(constant.AuthorizationHeader)
		if token == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": constant.InvalidToken,
			})
		}

		claims, err := jwt.VerifyToken(config, token)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		ctx.Locals(constant.UserContext, claims)
		return ctx.Next()
	}
}
