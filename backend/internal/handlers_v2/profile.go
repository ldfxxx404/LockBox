package handlersv2

import (
	"back/internal/models"
	"back/internal/services"
	"back/internal/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type ProfileHandlerV2 struct {
	ProfileServ     *services.ProfileService
	ProfileFileServ *services.FileService
}

func NewProfileHandlerV2(
	profileServ *services.ProfileService,
	profileFileServ *services.FileService,
) *ProfileHandlerV2 {
	return &ProfileHandlerV2{
		ProfileServ:     profileServ,
		ProfileFileServ: profileFileServ,
	}
}

// GetProfile godoc
// @Summary Get user profile v2
// @Description Returns the current user's profile and storage usage info
// @Tags profile v2
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.ProfileV2
// @Failure 500 {object} models.ErrorResponse
// @Router /v2/profile [get]
func (h *ProfileHandlerV2) GetProfile(c *fiber.Ctx) error {
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

	log.Info("GetProfileV2: profile retrieved successfully", "user_id", userID, "email", user.Email)
	return c.JSON(models.ProfileV2{
		Storage: &models.ProfileStorage{
			Used:  usedMB,
			Limit: limitMB,
		},
		Files: filenames,
		User: &models.UserV2{
			Name: user.Name,
		},
	})
}
