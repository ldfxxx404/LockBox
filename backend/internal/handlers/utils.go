package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ParamInt(c *fiber.Ctx, key string) (int, error) {
	return strconv.Atoi(c.Params(key))
}
