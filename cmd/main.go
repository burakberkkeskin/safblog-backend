package main

import (
	"fmt"
	"os"
	"safblog-backend/database"
	"safblog-backend/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()
	router.SetupRoutes(app)
	fmt.Println(os.Getenv("JWTSECRET"))
	fmt.Println("fiber app created")
	app.Listen(":8080")
}
