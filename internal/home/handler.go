package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHomeHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	api := h.router.Group("/api")
	api.Get("/", h.Home)
	api.Get("/error", h.GetError)
}

func (h *HomeHandler) Home(c *fiber.Ctx) error {
	h.customLogger.Info().Bool("isAdmin", true).Msg("Инфо")
	return c.SendString("Hello, World!")
}

func (h *HomeHandler) GetError(c *fiber.Ctx) error {
	return c.SendString("Error")
}
