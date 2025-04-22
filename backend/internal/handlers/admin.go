package handlers

import (
	"back/internal/models"
	"back/internal/services"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	AdminServ *services.AdminService
}

func NewAdminHandler(adminServ *services.AdminService) *AdminHandler {
	return &AdminHandler{AdminServ: adminServ}
}

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Returns a list of all registered users
// @Tags         admin
// @Security     BearerAuth
// @Produce      json
// @Success      200  {array}   models.User
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/users [get]
func (h *AdminHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.AdminServ.GetAllUsers()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{Massage: "model is no found", Error: err})
	}
	return c.JSON(users)
}

// UpdateStorageLimit godoc
// @Summary      Update user storage limit
// @Description  Sets a new storage limit for a specific user
// @Tags         admin
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body  body  models.UpdateStorage  true  "Storage limit info"
// @Success      200  {object}  models.SucessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Router       /admin/update_limit [post]
func (h *AdminHandler) UpdateStorageLimit(c *fiber.Ctx) error {
	var req models.UpdateStorage

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{Massage: "Invalid input", Error: err})
	}

	err := h.AdminServ.UpdateStorageLimit(req.UserID, req.NewLimit)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{Massage: "error", Error: err})
	}

	return c.JSON(models.SucessResponse{Massage: "Storage update limit succes"})
}

// MakeAdmin godoc
// @Summary      Grant admin rights
// @Description  Gives a user admin privileges
// @Tags         admin
// @Security     BearerAuth
// @Produce      json
// @Param        user_id  path  int  true  "User ID"
// @Success      200  {object}  models.SucessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/make_admin/{user_id} [post]
func (h *AdminHandler) MakeAdmin(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{Massage: "invalid user id", Error: err})
	}
	err = h.AdminServ.MakeAdmin(userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.ErrorResponse{Massage: "error", Error: err})
	}
	return c.JSON(models.SucessResponse{Massage: "User in now an admin"})
}

// RevokeAdmin godoc
// @Summary      Revoke admin rights
// @Description  Removes admin privileges from a user
// @Tags         admin
// @Security     BearerAuth
// @Produce      json
// @Param        user_id  path  int  true  "User ID"
// @Success      200  {object}  models.SucessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/revoke_admin/{user_id} [post]
func (h *AdminHandler) RevokeAdmin(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{Massage: "invalid user id", Error: err})
	}
	err = h.AdminServ.RevokeAdmin(userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.ErrorResponse{Massage: "error", Error: err})
	}
	return c.JSON(models.SucessResponse{Massage: "Admin rights revoked"})
}
