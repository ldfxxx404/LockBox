package handlers

import (
	"back/internal/models"
	"back/internal/services"
	"back/internal/utils"
	"net/http"

	"github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2"
)

type ProfileHandler struct {
	ProfileServ *services.ProfileService
}

func NewProfileHandler(
	profileServ *services.ProfileService,
) *ProfileHandler {
	return &ProfileHandler{
		ProfileServ: profileServ,
	}
}

// GetProfile godoc
// @Summary Get user profile
// @Description Returns the current user's profile and storage usage info
// @Tags profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Profile
// @Failure 500 {object} models.ErrorResponse
// @Router /profile [get]
func (h *ProfileHandler) GetProfile(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	user, usedMB, limitMB, err := h.ProfileServ.GetProfile(userID)
	if err != nil {
		log.Error("GetProfile: failed to retrieve profile", "user_id", userID, "error", err)
		return c.Status(http.StatusInternalServerError).
			JSON(models.ErrorResponse{
				Message: "Failed to get user profile",
				Error:   err.Error(),
			})
	}

	log.Info("GetProfile: profile retrieved successfully", "user_id", userID, "email", user.Email)
	return c.JSON(models.Profile{
		User: &models.ProfileUser{
			Id:    userID,
			Email: user.Email,
			Name:  user.Name,
		},
		Storage: &models.ProfileStorage{
			Used:  usedMB,
			Limit: limitMB,
		},
	})
}
