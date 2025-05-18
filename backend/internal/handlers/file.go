package handlers

import (
	"back/internal/models"
	"back/internal/services"
	"back/internal/utils"
	"net/http"
	"net/url"

	log "github.com/charmbracelet/log"

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
// @Success      200   {object}  models.FileUpload
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /upload [post]
func (h *FileHandler) Upload(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	usedMB, limitMB, _ := h.FileServ.GetStorageInfo(userID)

	fileHeader, err := c.FormFile("file")
	if err != nil {
		log.Error("handler: error of formFile upload", "err", err)
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{Message: "File is required", Error: err.Error()})
	}

	err = utils.CheckSize(usedMB, fileHeader.Size, limitMB)
	if err != nil {
		log.Error("handler: check file size", "err", err)
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{Message: "error", Error: err.Error()})
	}

	err = h.FileServ.UploadFile(userID, fileHeader)
	if err != nil {
		log.Error("handler: upload user file", "err", err)
		return c.Status(http.StatusInternalServerError).JSON(models.ErrorResponse{Message: "error", Error: err.Error()})
	}
	log.Info("success file upload")
	return c.JSON(models.FileUpload{Message: "File upload", FileName: fileHeader.Filename})
}

// ListFiles godoc
// @Summary      List user's files
// @Description  Returns a list of filenames and storage info for the authenticated user
// @Tags         file
// @Security     BearerAuth
// @Produce      json
// @Success      200   {object}  models.ListFile
// @Failure      500   {object}  models.ErrorResponse
// @Router       /storage [get]
func (h *FileHandler) ListFiles(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	files, err := h.FileServ.ListFiles(userID)
	if err != nil {
		log.Error("handler: list users files", "err", err)
		return c.Status(http.StatusInternalServerError).JSON(models.ErrorResponse{Message: "error", Error: err.Error()})
	}
	usedMB, limitMB, err := h.FileServ.GetStorageInfo(userID)
	if err != nil {
		log.Error("handler: get storage info", "err", err)
		return c.Status(http.StatusInternalServerError).JSON(models.ErrorResponse{Message: "error", Error: err.Error()})
	}
	var filenames []string
	for _, file := range files {
		filenames = append(filenames, file.Filename)
	}

	log.Info("success list files users")
	return c.JSON(models.ListFile{Files: filenames, Storage: &models.ListFileUsed{Used: usedMB, Limit: limitMB}})
}

// Download godoc
// @Summary      Download a file
// @Description  Downloads a specific file for the authenticated user
// @Tags         file
// @Security     BearerAuth
// @Produce      octet-stream
// @Param        filename  path  string  true  "File name"
// @Success      200
// @Failure      400   {object}  models.ErrorResponse
// @Failure      404   {object}  models.ErrorResponse
// @Router       /storage/{filename} [get]
func (h *FileHandler) Download(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	filename := c.Params("filename")

	decodedFilename, err := url.QueryUnescape(filename)
	if err != nil {
		log.Error("handler: get query of files", "err", err)
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{Message: "Invalid filename", Error: err.Error()})
	}

	filePath, err := h.FileServ.GetFilePath(userID, decodedFilename)
	if err != nil {
		log.Error("handler: get file path", "err", err)
		return c.Status(http.StatusNotFound).JSON(models.ErrorResponse{Message: "File not found", Error: err.Error()})
	}
	log.Info("success download file")
	return c.SendFile(filePath, false)
}

// Delete godoc
// @Summary      Delete a file
// @Description  Deletes a specific file belonging to the authenticated user
// @Tags         file
// @Security     BearerAuth
// @Param        filename  path  string  true  "File name"
// @Produce      json
// @Success      200   {object}  models.SuccessResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /delete/{filename} [delete]
func (h *FileHandler) Delete(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	filename := c.Params("filename")
	err := h.FileServ.DeleteFile(userID, filename)
	if err != nil {
		log.Error("handler: delete files", "err", err)
		return c.Status(http.StatusInternalServerError).JSON(models.ErrorResponse{Message: "error", Error: err.Error()})
	}
	return c.JSON(models.SuccessResponse{Message: "File delited"})
}
