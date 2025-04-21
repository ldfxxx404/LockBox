package handlers

import (
	"back/internal/models"
	"back/internal/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthServ *services.AuthService
}

func NewAuthHandler(authServ *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthServ: authServ}
}

// Register godoc
// @Summary      Register a new user
// @Description  Create an account with email and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.UserDTO  true  "User registration data"
// @Success      201   {object}  map[string]any
// @Failure      400   {object}  map[string]string
// @Router       /api/register [post]
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

// Login godoc
// @Summary      Log in user
// @Description  Authenticate user and return JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      models.LoginDTO  true  "User login credentials"
// @Success      200          {object}  map[string]interface{}
// @Failure      400          {object}  map[string]string
// @Failure      401          {object}  map[string]string
// @Router       /api/login [post]
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

// Logout godoc
// @Summary      Log out user
// @Description  Logout by removing token on the client side
// @Tags         auth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /api/logout [post]
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	// Для JWT logout осуществляется на клиентской стороне
	return c.JSON(fiber.Map{
		"message": "Logout successful",
	})
}
