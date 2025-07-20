package handlers

import (
	"back/internal/models"
	"back/internal/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
		log.Error("Failed to get all users", "error", err)
		return JSONError(c, http.StatusBadRequest, "Model not found", err)
	}
	return c.JSON(users)
}

// GetAllUsers godoc
// @Summary      Get all users who are admins
// @Description  Returns a list of all registered users
// @Tags         admin
// @Security     BearerAuth
// @Produce      json
// @Success      200  {array}   models.User
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/admins [get]
func (h *AdminHandler) GetAllAdminUsers(c *fiber.Ctx) error {
	admins, err := h.AdminServ.GetAllAdminUsers()
	if err != nil {
		log.Error("Failed to get all admin users", "error", err)
		return JSONError(c, http.StatusBadRequest, "Model not found", err)
	}
	return c.JSON(admins)
}

// UpdateStorageLimit godoc
// @Summary      Update user storage limit
// @Description  Sets a new storage limit for a specific user
// @Tags         admin
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body  body  models.UpdateStorage  true  "Storage limit info"
// @Success      200  {object}  models.SuccessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Router       /admin/update_limit [put]
func (h *AdminHandler) UpdateStorageLimit(c *fiber.Ctx) error {
	var req models.UpdateStorage
	if err := c.BodyParser(&req); err != nil {
		return JSONError(c, http.StatusBadRequest, "Invalid input", err)
	}

	if err := h.AdminServ.UpdateStorageLimit(req.UserID, req.NewLimit); err != nil {
		return JSONError(c, http.StatusBadRequest, "Error updating storage limit", err)
	}
	return JSONSuccess(c, "Storage limit updated successfully")
}

// MakeAdmin godoc
// @Summary      Grant admin rights
// @Description  Gives a user admin privileges
// @Tags         admin
// @Security     BearerAuth
// @Produce      json
// @Param        user_id  path  int  true  "User ID"
// @Success      200  {object}  models.SuccessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/make_admin/{user_id} [put]
func (h *AdminHandler) MakeAdmin(c *fiber.Ctx) error {
	userID, err := ParamInt(c, "user_id")
	if err != nil {
		return JSONError(c, http.StatusBadRequest, "Invalid user id", err)
	}

	if err := h.AdminServ.MakeAdmin(userID); err != nil {
		return JSONError(c, http.StatusInternalServerError, "Failed to grant admin rights", err)
	}
	return JSONSuccess(c, "User is now an admin")
}

// RevokeAdmin godoc
// @Summary      Revoke admin rights
// @Description  Removes admin privileges from a user
// @Tags         admin
// @Security     BearerAuth
// @Produce      json
// @Param        user_id  path  int  true  "User ID"
// @Success      200  {object}  models.SuccessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/revoke_admin/{user_id} [put]
func (h *AdminHandler) RevokeAdmin(c *fiber.Ctx) error {
	userID, err := ParamInt(c, "user_id")
	if err != nil {
		return JSONError(c, http.StatusBadRequest, "Invalid user id", err)
	}

	if err := h.AdminServ.RevokeAdmin(userID); err != nil {
		return JSONError(c, http.StatusInternalServerError, "Failed to revoke admin rights", err)
	}
	return JSONSuccess(c, "Admin rights revoked")
}
