package services

import (
	ResponseModels "safblog-backend/models/response_models"

	"github.com/gofiber/fiber/v2"
)

func CreateResponse(message string, data fiber.Map, error string) ResponseModels.Response {
	var response = ResponseModels.Response{
		Message: message,
		Data:    data,
		Error:   error,
	}
	return response
}

func SuccessResponse(message string, data fiber.Map) ResponseModels.Response {
	var response = CreateResponse(message, data, "")
	return response
}

func ErrorResponse(message string, error string) ResponseModels.Response {
	var response = CreateResponse(message, nil, error)
	return response
}
