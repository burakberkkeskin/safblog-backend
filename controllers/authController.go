package controllers

import (
	"fmt"
	"safblog-backend/models"
	"safblog-backend/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterController(c *fiber.Ctx) error {
	fmt.Println("Register Controller")
	var registerUser models.RegisterModel
	err := c.BodyParser(&registerUser)
	if err != nil {
		return err
	}
	response := services.CreateUser(registerUser)
	c.JSON(response)
	return nil
}

func LoginController(c *fiber.Ctx) error {
	fmt.Println("Login Controller")
	var loginUser models.LoginUser
	err := c.BodyParser(&loginUser)
	if err != nil {
		return err
	}
	response, err := services.LoginUser(loginUser)
	if err != nil {
		if err.Error() == "user not found" {
			c.Status(404).JSON(response)
		} else if err.Error() == "credentials are not valid" {
			c.Status(401).JSON(response)
		} else {
			return c.Status(500).JSON(models.Response{Status: "error", Error: "Something happened while logging in."})
		}
	}
	c.JSON(response)
	return nil
}
