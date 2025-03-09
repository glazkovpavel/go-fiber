package main

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go/go-fiber/config"
	"go/go-fiber/internal/home"
	"go/go-fiber/internal/vacancy"
	"go/go-fiber/pkg/logger"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New()) // для того чтобы не падало приложение при панике
	app.Static("/public", "./public")
	home.NewHomeHandler(app, customLogger)
	vacancy.NewHandler(app, customLogger)
	app.Listen(":3000")
}
