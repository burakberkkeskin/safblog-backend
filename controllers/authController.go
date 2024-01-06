package controllers

import (
	"fmt"
	"safblog-backend/models"
	"safblog-backend/services"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
)

func RegisterController(c *fiber.Ctx) error {
	fmt.Println("Register Controller")
	var registerUser models.RegisterModel
	err := c.BodyParser(&registerUser)
	if err != nil {
		return err
	}
	response, err := services.CreateUser(registerUser)
	if err != nil {
		switch err.Error() {
		case "error while hashing the password":
			c.Status(500).JSON(response)
		case "email already in use":
			c.Status(409).JSON(response)
		default:
			c.Status(400).JSON(response)
		}
	}
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
			return c.Status(500).JSON(models.Response{Message: "error", Error: "Something happened while logging in."})
		}
	}
	c.JSON(response)
	return nil
}

func WhoAmIController(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userEmail := claims["email"].(string)
	userId := claims["id"].(string)
	c.SendString("Welcome 👋" + userEmail + " " + userId)
	return nil
}
