package handlers

import (
	"back/internal/models"

	"github.com/gofiber/fiber/v2"
)

func JSONError(c *fiber.Ctx, status int, msg string, err error) error {
	return c.Status(status).JSON(models.ErrorResponse{
		Message: msg,
		Error:   err.Error(),
	})
}

func JSONSuccess(c *fiber.Ctx, msg string) error {
	return c.JSON(models.SuccessResponse{
		Message: msg,
	})
}
