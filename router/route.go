package router

import (
	"fmt"
	"os"
	"safblog-backend/controllers"
	"safblog-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// grouping
	fmt.Println("Setting up routes")
	fmt.Println("Getting routes ready.")
	jwtMiddleware := middlewares.NewAuthMiddleware(os.Getenv("JWTSECRET"))

	api := app.Group("/api")
	version := api.Group("/v1")

	// auth routes
	auth := version.Group("/auth")
	auth.Post("/register", controllers.RegisterController)
	auth.Post("/login", controllers.LoginController)
	auth.Get("/whoami", jwtMiddleware, controllers.WhoAmIController)
	// other routes

}
