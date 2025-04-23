package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetUserID(c *fiber.Ctx) int {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return int(claims["user_id"].(float64))
}

func CheckSize(used, size int64, limit int) error {
	sizeMB := size / (1024 * 1024)
	if used+sizeMB > int64(limit) {
		return errors.New("file size is too large")
	}
	return nil
}
