package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"go/go-fiber/internal/vacancy"
	"go/go-fiber/pkg/tadapter"
	"go/go-fiber/views"
	"math"
	"net/http"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *vacancy.VacancyRepository
}

type User struct {
	Id   int
	Name string
}

func NewHomeHandler(router fiber.Router, customLogger *zerolog.Logger, repository *vacancy.VacancyRepository) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
	}
	h.router.Get("/", h.Home)
	h.router.Get("/404", h.GetError)
}

func (h *HomeHandler) Home(c *fiber.Ctx) error {
	PAGE_ITEMS := 2
	page := c.QueryInt("page", 1)
	count := h.repository.CountAll()
	vacancies, err := h.repository.GetAll(PAGE_ITEMS, (page-1)*PAGE_ITEMS)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	h.customLogger.Info().Bool("isAdmin", true).Msg("Инфо")
	component := views.Main(vacancies, int(math.Ceil(float64(count/PAGE_ITEMS))), page)
	return tadapter.Render(c, component, http.StatusOK)
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
