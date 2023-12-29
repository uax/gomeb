package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/uax/gomeb/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		AppName: os.Getenv("APP_NAME"),
	})
	router.RegisterRouter(app)

	app.Listen(os.Getenv("APP_LISTEN"))
}
