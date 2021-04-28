package main

import (
	"os"

	"example.com/compA"
	"example.com/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// create fiber instance
	app := fiber.New(fiber.Config{
		Prefork: os.Getenv("GO_ENV") == "production",
	})

	// health check endpoint
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// initialize DB and run migration on all models
	// https://gorm.io/docs/migration.html
	db := database.Init()
	db.AutoMigrate(&compA.CompA{} /*, &compB.CompB */)

	// register routes
	api := app.Group("/api")
	compA.RegisterRoute(api)
	// compB.RegisterRoute(api)

	// run app
	app.Listen(":3000")
}
