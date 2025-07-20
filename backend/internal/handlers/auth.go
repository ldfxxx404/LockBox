package handlers

import (
	"back/internal/models"
	"back/internal/services"
	"net/http"

	"github.com/gofiber/fiber/v2/log"

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
// @Param        user  body      models.RegisterDTO  true  "User registration data"
// @Success      201   {object}  models.RegisterMessage
// @Failure      400   {object}  models.ErrorResponse
// @Router       /register [post]

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var dto models.RegisterDTO
	if err := c.BodyParser(&dto); err != nil {
		return JSONError(c, http.StatusBadRequest, "Invalid input", err)
	}

	user, err := h.AuthServ.Register(dto)
	if err != nil {
		return JSONError(c, http.StatusBadRequest, "Registration error", err)
	}

	log.Info("User registered", "user_id", user.ID, "email", user.Email)
	return c.Status(http.StatusCreated).JSON(models.RegisterMessage{
		Message: "Registration successful",
		User: &models.MessageDTO{
			Id:    user.ID,
			Email: user.Email,
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
// @Success      200          {object}  models.LoginMessage
// @Failure      400          {object}  models.ErrorResponse
// @Failure      401          {object}  models.ErrorResponse
// @Router       /login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var dto models.LoginDTO
	if err := c.BodyParser(&dto); err != nil {
		return JSONError(c, http.StatusBadRequest, "Invalid input", err)
	}

	token, user, err := h.AuthServ.Login(dto)
	if err != nil {
		return JSONError(c, http.StatusUnauthorized, "Authentication failed", err)
	}

	log.Info("Login successful", "user_id", user.ID, "email", user.Email)
	return c.JSON(models.LoginMessage{
		Message: "Login successful",
		User: &models.MessageDTO{
			Id:    user.ID,
			Email: user.Email,
		},
		Token: token,
	})
}

// Logout godoc
// @Summary      Log out user
// @Description  Logout by removing token on the client side
// @Tags         auth
// @Produce      json
// @Success      200  {object}  models.SuccessResponse
// @Router       /logout [post]
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	log.Warn("Logout called — client should remove token")
	return JSONSuccess(c, "Logout successful (client-side token removal required)")
}