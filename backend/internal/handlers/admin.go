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
		log.Debug(err)
		log.Error("handler: error of getallsers", "err", err)
		return c.Status(http.StatusBadRequest).
			JSON(models.ErrorResponse{
				Message: "model is no found",
				Error:   err.Error(),
			})
	}
	log.Info("success get all users")
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
// @Success      200  {object}  models.SuccessResponse
// @Failure      400  {object}  models.ErrorResponse
// @Router       /admin/update_limit [put]
func (h *AdminHandler) UpdateStorageLimit(c *fiber.Ctx) error {
	var req models.UpdateStorage

	if err := c.BodyParser(&req); err != nil {
		log.Error("handler: body parcer error", "err", err)
		return c.Status(http.StatusBadRequest).
			JSON(models.ErrorResponse{
				Message: "Invalid input",
				Error:   err.Error(),
			})
	}

	err := h.AdminServ.UpdateStorageLimit(req.UserID, req.NewLimit)
	if err != nil {
		log.Error("handler: update storage limit", "err", err)
		return c.Status(http.StatusBadRequest).
			JSON(models.ErrorResponse{
				Message: "error",
				Error:   err.Error(),
			})
	}
	log.Info("success update storage limit")
	return c.JSON(models.SuccessResponse{Message: "Storage update limit succes"})
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
		log.Error("handler: get params user Ids", "err", err)
		return c.Status(http.StatusBadRequest).
			JSON(models.ErrorResponse{
				Message: "invalid user id",
				Error:   err.Error(),
			})
	}
	err = h.AdminServ.MakeAdmin(userID)
	if err != nil {
		log.Error("handler: make admin", "err", err)
		return c.Status(http.StatusInternalServerError).
			JSON(models.ErrorResponse{
				Message: "error",
				Error:   err.Error(),
			})
	}
	log.Info("success user make admin")
	return c.JSON(models.SuccessResponse{Message: "User in now an admin"})
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
		log.Error("handler: revoke admin", "err", err)
		return c.Status(http.StatusBadRequest).
			JSON(models.ErrorResponse{
				Message: "invalid user id",
				Error:   err.Error(),
			})
	}
	err = h.AdminServ.RevokeAdmin(userID)
	if err != nil {
		log.Error("handler: revoke admin", "err", err)
		return c.Status(http.StatusInternalServerError).
			JSON(models.ErrorResponse{
				Message: "error",
				Error:   err.Error(),
			})
	}
	log.Info("success Revoke admin")
	return c.JSON(models.SuccessResponse{Message: "Admin rights revoked"})
}
