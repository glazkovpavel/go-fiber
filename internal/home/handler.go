package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

type User struct {
	Id   int
	Name string
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

	users := []User{
		{Id: 1, Name: "Petr"},
		{Id: 2, Name: "Cat"},
	}

	names := []string{"Антон", "Вася"}
	data := struct {
		Users []User
		Names []string
	}{Names: names, Users: users}
	return c.Render("slice", data)
}

func (h *HomeHandler) GetError(c *fiber.Ctx) error {
	//tmpl := template.Must(template.ParseFiles("./html/page.html"))

	//
	//var tpl bytes.Buffer
	//err := tmpl.Execute(&tpl, data)
	//if err != nil {
	//	return fiber.NewError(fiber.StatusBadRequest, err.Error())
	//}
	//c.Set(fiber.HeaderContentType, "text/html; charset=utf-8")

	data := struct {
		Count   int
		IsAdmin bool
		CanUse  bool
	}{Count: 1, IsAdmin: true, CanUse: true}

	return c.Render("page", data)
}
