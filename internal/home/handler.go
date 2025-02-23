package home

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {
	router fiber.Router
}

func NewHomeHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	h.router.Get("/", h.Home)
}

func (h *HomeHandler) Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
