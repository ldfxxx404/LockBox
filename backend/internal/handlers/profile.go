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
	ProfileServ     *services.ProfileService
	ProfileFileServ *services.FileService
}

func NewProfileHandler(
	profileServ *services.ProfileService,
	profileFileServ *services.FileService,
) *ProfileHandler {
	return &ProfileHandler{
		ProfileServ:     profileServ,
		ProfileFileServ: profileFileServ,
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

// GetProfile godoc
// @Summary Get user profile v2
// @Description Returns the current user's profile and storage usage info
// @Tags profile v2
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.ProfileStorage1
// @Failure 500 {object} models.ErrorResponse
// @Router /v2/profile [get]
func (h *ProfileHandler) GetV2profile(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	user, usedMB, limitMB, err := h.ProfileServ.GetProfile(userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(models.ErrorResponse{
				Message: "Failed to get user profile",
				Error:   err.Error(),
			})
	}
	files, err := h.ProfileFileServ.ListFiles(userID)
	if err != nil {
		log.Error("ListFiles: failed to list files", "user_id", userID, "err", err)
		return c.Status(http.StatusInternalServerError).
			JSON(models.ErrorResponse{
				Message: "Failed to list files",
				Error:   err.Error(),
			})
	}

	var filenames []string
	for _, file := range files {
		filenames = append(filenames, file.Filename)
	}

	return c.JSON(models.ProfileStorage1{
		Storage: &models.ProfileStorage{
			Used:  usedMB,
			Limit: limitMB,
		},
		Files: filenames,
		User: &models.ProfileUser1{
			Name: user.Name,
		},
	})
}
