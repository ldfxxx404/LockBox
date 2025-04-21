package handlers

import (
	"back/internal/models"
	"back/internal/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthHandler struct {
	AuthServ *services.AuthService
}

func NewAuthHandler(authServ *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthServ: authServ}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var dto models.UserDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	user, err := h.AuthServ.Register(dto)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Registration successful",
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var dto models.LoginDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	token, user, err := h.AuthServ.Login(dto)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
		},
		"token": token,
	})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	// Для JWT logout осуществляется на клиентской стороне
	return c.JSON(fiber.Map{
		"message": "Logout successful",
	})
}
