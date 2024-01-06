package services

import (
	"safblog-backend/models"

	"github.com/gofiber/fiber/v2"
)

func CreateResponse(message string, data fiber.Map, error string) models.Response {
	var response = models.Response{
		Message: message,
		Data:    data,
		Error:   error,
	}
	return response
}
