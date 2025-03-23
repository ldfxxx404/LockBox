package handlers

import (
	"back/internal/services"
	"back/internal/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ProfileHandler struct {
	ProfileServ *services.ProfileService
}

func NewProfileHandler(profileServ *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{ProfileServ: profileServ}
}

func (h *ProfileHandler) GetProfile(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	user, usedMB, limitMB, err := h.ProfileServ.GetProfile(userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
		},
		"storage": fiber.Map{
			"used":  usedMB,
			"limit": limitMB,
		},
	})
}
