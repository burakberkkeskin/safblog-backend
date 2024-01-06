package main

import (
	"fmt"
	"safblog-backend/database"
	"safblog-backend/router"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))
	fmt.Println("fiber app created")
	router.SetupRoutes(app)
	app.Listen(":8080")
}
