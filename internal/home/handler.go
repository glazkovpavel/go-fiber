package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
	"go/go-fiber/internal/vacancy"
	"go/go-fiber/pkg/tadapter"
	"go/go-fiber/views"
	"go/go-fiber/views/components"
	"math"
	"net/http"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *vacancy.VacancyRepository
	store        *session.Store
}

type User struct {
	Id   int
	Name string
}

func NewHomeHandler(router fiber.Router, customLogger *zerolog.Logger, repository *vacancy.VacancyRepository, store *session.Store) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
		store:        store,
	}
	h.router.Get("/", h.Home)
	h.router.Get("/login", h.Login)
	h.router.Post("/api/login", h.apiLogin)
	h.router.Get("/api/logout", h.apiLogout)
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

func (h *HomeHandler) apiLogout(c *fiber.Ctx) error {
	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Delete("email")
	if err := sess.Save(); err != nil {
		panic(err)
	}
	c.Response().Header.Add("Hx-Redirect", "/")
	return c.Redirect("/", http.StatusOK)
}

func (h *HomeHandler) apiLogin(c *fiber.Ctx) error {
	form := LoinForm{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	if form.Email == "a@a.ru" || form.Password == "1" {
		sess, err := h.store.Get(c)
		if err != nil {
			panic(err)
		}
		sess.Set("email", form.Email)
		if err := sess.Save(); err != nil {
			panic(err)
		}
		c.Response().Header.Add("Hx-Redirect", "/")
		return c.Redirect("/", http.StatusOK)
	}
	component := components.Notification("Неверный логин или пароль", components.NotificationFail)
	return tadapter.Render(c, component, http.StatusBadRequest)
}

func (h *HomeHandler) Login(c *fiber.Ctx) error {
	component := views.Login()
	return tadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) GetError(c *fiber.Ctx) error {
	data := struct {
		Count   int
		IsAdmin bool
		CanUse  bool
	}{Count: 1, IsAdmin: true, CanUse: true}

	return c.Render("page", data)
}
