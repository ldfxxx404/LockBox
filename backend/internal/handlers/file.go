package handlers

import (
	"back/internal/models"
	"back/internal/services"
	"back/internal/utils"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2/log"

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
		log.Error("Upload: missing file in form data", "user_id", userID, "err", err)
		return c.Status(http.StatusBadRequest).
			JSON(models.ErrorResponse{
				Message: "File is required",
				Error:   err.Error(),
			})
	}

	err = utils.CheckSize(usedMB, fileHeader.Size, limitMB)
	if err != nil {
		log.Error("Upload: file size exceeds limit", "user_id", userID, "file_size", fileHeader.Size, "used_mb", usedMB, "limit_mb", limitMB, "err", err)
		return c.Status(http.StatusBadRequest).
			JSON(models.ErrorResponse{
				Message: "File size exceeds storage limit",
				Error:   err.Error(),
			})
	}

	err = h.FileServ.UploadFile(userID, fileHeader)
	if err != nil {
		log.Error("Upload: failed to upload file", "user_id", userID, "file_name", fileHeader.Filename, "err", err)
		return c.Status(http.StatusInternalServerError).
			JSON(models.ErrorResponse{
				Message: "Failed to upload file",
				Error:   err.Error(),
			})
	}
	log.Info("Upload: file uploaded successfully", "user_id", userID, "file_name", fileHeader.Filename)
	return c.JSON(models.FileUpload{
		Message:  "File uploaded successfully",
		FileName: fileHeader.Filename,
	})
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
		log.Error("ListFiles: failed to list files", "user_id", userID, "err", err)
		return c.Status(http.StatusInternalServerError).
			JSON(models.ErrorResponse{
				Message: "Failed to list files",
				Error:   err.Error(),
			})
	}
	usedMB, limitMB, err := h.FileServ.GetStorageInfo(userID)
	if err != nil {
		log.Error("ListFiles: failed to get storage info", "user_id", userID, "err", err)
		return c.Status(http.StatusInternalServerError).
			JSON(models.ErrorResponse{
				Message: "Failed to get storage info",
				Error:   err.Error(),
			})
	}

	var filenames []string
	for _, file := range files {
		filenames = append(filenames, file.Filename)
	}

	log.Info("ListFiles: retrieved file list", "user_id", userID, "files_count", len(filenames))
	return c.JSON(models.ListFile{
		Files: filenames,
		Storage: &models.ListFileUsed{
			Used:  usedMB,
			Limit: limitMB,
		},
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
// @Failure      400   {object}  models.ErrorResponse
// @Failure      404   {object}  models.ErrorResponse
// @Router       /storage/{filename} [get]
func (h *FileHandler) Download(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	filename := c.Params("filename")

	decodedFilename, err := url.QueryUnescape(filename)
	if err != nil {
		log.Error("Download: invalid filename parameter", "user_id", userID, "filename", filename, "err", err)
		return c.Status(http.StatusBadRequest).
			JSON(models.ErrorResponse{
				Message: "Invalid filename",
				Error:   err.Error(),
			})
	}

	data, err := h.FileServ.GetFile(userID, decodedFilename)
	if err != nil {
		log.Error("Download: file not found or inaccessible", "user_id", userID, "filename", decodedFilename, "err", err)
		return c.Status(http.StatusNotFound).
			JSON(models.ErrorResponse{
				Message: "File not found",
				Error:   err.Error(),
			})
	}

	log.Info("Download: file served successfully", "user_id", userID, "filename", decodedFilename)

	c.Set("Content-Type", "application/octet-stream")
	c.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, decodedFilename))

	return c.Send(data)
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
	filename, _ := url.QueryUnescape(c.Params("filename"))
	err := h.FileServ.DeleteFile(userID, filename)
	if err != nil {
		log.Error("Delete: failed to delete file", "user_id", userID, "filename", filename, "err", err)
		return c.Status(http.StatusInternalServerError).
			JSON(models.ErrorResponse{
				Message: "Failed to delete file",
				Error:   err.Error(),
			})
	}
	log.Info("Delete: file deleted successfully", "user_id", userID, "filename", filename)
	return c.JSON(models.SuccessResponse{Message: "File deleted"})
}
