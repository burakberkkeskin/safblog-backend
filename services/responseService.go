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

func SuccessResponse(message string, data fiber.Map) models.Response {
	var response = CreateResponse(message, data, "")
	return response
}

func ErrorResponse(message string, error string) models.Response {
	var response = CreateResponse(message, nil, error)
	return response
}
