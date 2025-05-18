package middleware

import (
	"back/config"
	"back/internal/models"

	"back/internal/models"

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
    return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
        Message: "Unauthorized",
        Error:   err.Error(),
    })
}

func unauthorizedResponse(c *fiber.Ctx) error {
    return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
        Message: "Unauthorized",
        Error:   "access denied",
    })
}

func AdminRequired() fiber.Handler {
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
        isAdmin, ok := claims["is_admin"].(bool)
        if !ok || !isAdmin {
            return unauthorizedResponse(c)
        }
        return c.Next()
    }
}
