package handlers

import (
	"back/internal/models"
	"back/internal/services"
	"net/http"
	"strconv"

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
		return c.Status(http.StatusBadRequest).
			JSON(models.ErrorResponse{
				Message: "Model not found",
				Error:   err.Error(),
			})
	}
	log.Info("Successfully retrieved all users", "count", len(users))
	return c.JSON(users)
}

// GetAllUsers godoc
// @Summary      Get all users who is admins
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
		log.Error("Failed to get all users", "error", err)
		return c.Status(http.StatusBadRequest).
			JSON(models.ErrorResponse{
				Message: "Model not found",
				Error:   err.Error(),
			})
	}
	log.Info("Successfully retrieved all users", "count", len(admins))
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
		log.Error("Body parsing failed", "error", err)
		return c.Status(http.StatusBadRequest).
			JSON(models.ErrorResponse{
				Message: "Invalid input",
				Error:   err.Error(),
			})
	}

	err := h.AdminServ.UpdateStorageLimit(req.UserID, req.NewLimit)
	if err != nil {
		log.Error("Failed to update storage limit", "user_id", req.UserID, "new_limit", req.NewLimit, "error", err)
		return c.Status(http.StatusBadRequest).
			JSON(models.ErrorResponse{
				Message: "Error updating storage limit",
				Error:   err.Error(),
			})
	}
	log.Info("Storage limit updated", "user_id", req.UserID, "new_limit", req.NewLimit)
	return c.JSON(models.SuccessResponse{Message: "Storage limit updated successfully"})
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
	userID, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		log.Error("Invalid user ID parameter", "param", c.Params("user_id"), "error", err)
		return c.Status(http.StatusBadRequest).
			JSON(models.ErrorResponse{
				Message: "Invalid user id",
				Error:   err.Error(),
			})
	}
	err = h.AdminServ.MakeAdmin(userID)
	if err != nil {
		log.Error("Failed to grant admin rights", "user_id", userID, "error", err)
		return c.Status(http.StatusInternalServerError).
			JSON(models.ErrorResponse{
				Message: "Failed to grant admin rights",
				Error:   err.Error(),
			})
	}
	log.Info("Granted admin rights", "user_id", userID)
	return c.JSON(models.SuccessResponse{Message: "User is now an admin"})
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
	userID, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		log.Error("Invalid user ID parameter", "param", c.Params("user_id"), "error", err)
		return c.Status(http.StatusBadRequest).
			JSON(models.ErrorResponse{
				Message: "Invalid user id",
				Error:   err.Error(),
			})
	}
	err = h.AdminServ.RevokeAdmin(userID)
	if err != nil {
		log.Error("Failed to revoke admin rights", "user_id", userID, "error", err)
		return c.Status(http.StatusInternalServerError).
			JSON(models.ErrorResponse{
				Message: "Failed to revoke admin rights",
				Error:   err.Error(),
			})
	}
	log.Info("Revoked admin rights", "user_id", userID)
	return c.JSON(models.SuccessResponse{Message: "Admin rights revoked"})
}
