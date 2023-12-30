package main

import (
	"fmt"
	"safblog-backend/database"
	"safblog-backend/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()
	fmt.Println("fiber app created")
	router.SetupRoutes(app)
	app.Listen(":8080")
}
