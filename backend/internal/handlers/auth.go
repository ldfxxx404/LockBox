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
// @Success      201   {object}  models.RegisterMassage
// @Failure      400   {object}  models.ErrorResponse
// @Router       /api/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var dto models.UserDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{Massage: "invalid input", Error: err})
	}
	user, err := h.AuthServ.Register(dto)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{Massage: "error", Error: err})
	}

	return c.Status(http.StatusCreated).JSON(models.RegisterMassage{Message: "Registration successful", User: &models.MassageDTO{Id: user.ID, Email: user.Email}})
}

// Login godoc
// @Summary      Log in user
// @Description  Authenticate user and return JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      models.LoginDTO  true  "User login credentials"
// @Success      200          {object}  models.LoginMassage
// @Failure      400          {object}  models.ErrorResponse
// @Failure      401          {object}  models.ErrorResponse
// @Router       /api/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var dto models.LoginDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{Massage: "error", Error: err})
	}
	token, user, err := h.AuthServ.Login(dto)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(models.ErrorResponse{Massage: "error", Error: err})
	}

	return c.JSON(models.LoginMassage{Message: "Login successful", User: &models.MassageDTO{Id: user.ID, Email: user.Email}, Token: token})
}

// Logout godoc
// @Summary      Log out user
// @Description  Logout by removing token on the client side
// @Tags         auth
// @Produce      json
// @Success      200  {object}  models.SucessResponse
// @Router       /api/logout [post]
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	// Для JWT logout осуществляется на клиентской стороне
	return c.JSON(models.SucessResponse{Massage: "Logout succesfull"})
}
