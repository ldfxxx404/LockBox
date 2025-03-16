package handlers

import (
	"back/internal/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	AdminServ *services.AdminService
}

func NewAdminHandler(adminServ *services.AdminService) *AdminHandler {
	return &AdminHandler{AdminServ: adminServ}
}

func (h *AdminHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.AdminServ.GetAllUsers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (h *AdminHandler) UpdateStorageLimit(c *fiber.Ctx) error {
	var req struct {
		UserID   int `json:"user_id"`
		NewLimit int `json:"new_limit"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	err := h.AdminServ.UpdateStorageLimit(req.UserID, req.NewLimit)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Storage limit updated"})
}

func (h *AdminHandler) MakeAdmin(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	err = h.AdminServ.MakeAdmin(userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "User is now an admin"})
}

func (h *AdminHandler) RevokeAdmin(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	err = h.AdminServ.RevokeAdmin(userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Admin rights revoked"})
}
