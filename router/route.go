package router

import (
	"fmt"
	"safblog-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// grouping
	fmt.Println("Setting up routes")
	fmt.Println("Getting routes ready.")

	api := app.Group("/api")

	// routes
	auth := api.Group("/auth")
	auth.Post("/register", controllers.RegisterController)
	auth.Post("/login", controllers.LoginController)
}
