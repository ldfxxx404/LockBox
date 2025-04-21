package handlers

import (
	"back/internal/services"
	"back/internal/utils"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type FileHandler struct {
	FileServ *services.FileService
}

func NewFileHandler(fileServ *services.FileService) *FileHandler {
	return &FileHandler{FileServ: fileServ}
}

// Upload godoc
// @Summary      Upload a file
// @Description  Upload a file for the authenticated user
// @Tags         file
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "File to upload"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /upload [post]
func (h *FileHandler) Upload(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "File is required"})
	}
	err = h.FileServ.UploadFile(userID, fileHeader)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "File uploaded", "filename": fileHeader.Filename})
}

// ListFiles godoc
// @Summary      List user's files
// @Description  Returns a list of filenames and storage info for the authenticated user
// @Tags         file
// @Security     BearerAuth
// @Produce      json
// @Success      200   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]string
// @Router       /storage [get]
func (h *FileHandler) ListFiles(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	files, err := h.FileServ.ListFiles(userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	usedMB, limitMB, err := h.FileServ.GetStorageInfo(userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var filenames []string
	for _, file := range files {
		filenames = append(filenames, file.Filename)
	}
	return c.JSON(fiber.Map{
		"files":   filenames,
		"storage": fiber.Map{"used": usedMB, "limit": limitMB},
	})
}

// Download godoc
// @Summary      Download a file
// @Description  Downloads a specific file for the authenticated user
// @Tags         file
// @Security     BearerAuth
// @Produce      octet-stream
// @Param        filename  path  string  true  "File name"
// @Success      200
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /storage/{filename} [get]
func (h *FileHandler) Download(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	filename := c.Params("filename")

	decodedFilename, err := url.QueryUnescape(filename)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid filename"})
	}

	filePath, err := h.FileServ.GetFilePath(userID, decodedFilename)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendFile(filePath, false)
}

// Delete godoc
// @Summary      Delete a file
// @Description  Deletes a specific file belonging to the authenticated user
// @Tags         file
// @Security     BearerAuth
// @Param        filename  path  string  true  "File name"
// @Produce      json
// @Success      200   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /delete/{filename} [delete]
func (h *FileHandler) Delete(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	filename := c.Params("filename")
	err := h.FileServ.DeleteFile(userID, filename)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "File deleted"})
}
