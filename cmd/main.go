package main

import (
	"safblog-backend/database"
	"safblog-backend/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(":8080")
}
