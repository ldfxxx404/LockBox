package middleware

import (
	"back/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func AuthRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.JWTSecret),
		ContextKey:   "user",
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

func AdminRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userToken := c.Locals("user")
		if userToken == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
		token, ok := userToken.(*jwt.Token)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
		if isAdmin, ok := claims["is_admin"].(bool); ok && isAdmin {
			return c.Next()
		}
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Admin access required",
		})
	}
}
