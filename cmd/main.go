package main

import (
	"safblog-backend/database"
	"safblog-backend/router"
	"safblog-backend/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	router.SetupRoutes(app)
	_, err := services.CreateRootUser()
	if err != nil {

	}
	app.Listen(":8080")
}
