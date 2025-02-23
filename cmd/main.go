package main

import (
	"github.com/gofiber/fiber/v2"
	"go/go-fiber/internal/home"
)

func main() {
	app := fiber.New()

	home.NewHomeHandler(app)
	app.Listen(":3000")
}
