package user

import (
	"errors"
	"github.com/faizinahsan/academic-system/internal/controller/http/v1/response"
	customError "github.com/faizinahsan/academic-system/pkg/custom-error"
	"github.com/gofiber/fiber/v2"
)

func errorResponse(ctx *fiber.Ctx, err error) error {
	var resErr error
	switch {
	case errors.Is(err, fiber.ErrUnauthorized):
		resErr = ctx.Status(fiber.StatusUnauthorized).JSON(response.Error{Error: "unauthorized"})
	case errors.Is(err, customError.StatusNotActive):
		resErr = ctx.Status(fiber.StatusUnauthorized).JSON(response.Error{Error: customError.StatusNotActive.Error()})
	default:
		resErr = ctx.Status(fiber.StatusInternalServerError).JSON(response.Error{Error: "internal server error"})
	}
	return resErr
}
