package middleware

import (
	"back/config"
	"back/internal/models"
	"back/internal/services"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func AdminRequired(adminService *services.AdminService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userToken := c.Locals("user")
		if userToken == nil {
			return unauthorizedResponse(c)
		}
		token, ok := userToken.(*jwt.Token)
		if !ok {
			return unauthorizedResponse(c)
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return unauthorizedResponse(c)
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return unauthorizedResponse(c)
		}
		userID := int(userIDFloat)

		isAdmin, err := adminService.IsAdmin(userID)
		if err != nil {
			return unauthorizedResponse(c)
		}
		if !isAdmin {
			return unauthorizedResponse(c)
		}
		return c.Next()
	}
}

func unauthorizedResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
		Message: "Unauthorized",
		Error:   "access denied",
	})
}

func AuthRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.JWTSecret),
		ContextKey:   "user",
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
		Message: "Unauthorized",
		Error:   err.Error(),
	})
}
