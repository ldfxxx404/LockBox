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

func (h *FileHandler) Delete(c *fiber.Ctx) error {
	userID := utils.GetUserID(c)
	filename := c.Params("filename")
	err := h.FileServ.DeleteFile(userID, filename)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "File deleted"})
}
