package main

import (
	"fmt"
	"safblog-backend/database"
	"safblog-backend/router"
	"safblog-backend/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	router.SetupRoutes(app)
	err := services.CreateRootUser()
	if err != nil {
		fmt.Println("Error while creating the root user: ", err)
	}
	app.Listen(":8080")
}
