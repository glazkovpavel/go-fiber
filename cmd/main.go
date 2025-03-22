package main

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
	"go/go-fiber/config"
	"go/go-fiber/internal/home"
	"go/go-fiber/internal/vacancy"
	"go/go-fiber/pkg/database"
	"go/go-fiber/pkg/logger"
	"go/go-fiber/pkg/middleware"
	"time"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	dbConfig := config.NewDatabaseConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New()) // для того чтобы не падало приложение при панике
	app.Static("/public", "./public")
	dbpool := database.CreateDbPool(dbConfig, customLogger)
	defer dbpool.Close()
	storage := postgres.New(postgres.Config{
		DB:         dbpool,
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})
	store := session.New(session.Config{
		Storage: storage,
	})
	app.Use(middleware.AuthMiddleware(store))
	// Repositories
	vacancyRepo := vacancy.NewVacancyRepository(dbpool, customLogger)
	// Handler
	home.NewHomeHandler(app, customLogger, vacancyRepo, store)
	vacancy.NewHandler(app, customLogger, vacancyRepo)
	app.Listen(":3000")
}
