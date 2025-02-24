package main

import (
	"github.com/gofiber/fiber/v2"
	"go/go-fiber/config"
	"go/go-fiber/internal/home"
	"log"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)
	app := fiber.New()

	home.NewHomeHandler(app)
	app.Listen(":3000")
}
