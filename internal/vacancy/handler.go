package vacancy

import (
	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"go/go-fiber/pkg/tadapter"
	"go/go-fiber/pkg/validator"
	"go/go-fiber/views/components"
	"net/http"
)

type VacancyHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *VacancyRepository
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *VacancyRepository) {
	h := &VacancyHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	form := VacancyCreateForm{
		Email:    c.FormValue("email"),
		Location: c.FormValue("location"),
		Salary:   c.FormValue("salary"),
		Role:     c.FormValue("role"),
		Company:  c.FormValue("company"),
		Type:     c.FormValue("type"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{
			Name: "Email", Field: form.Email, Message: "Email не верен или не задан",
		},
		&validators.StringIsPresent{
			Name: "Location", Field: form.Location, Message: "Расположение не задано",
		},
		&validators.StringIsPresent{
			Name: "Type", Field: form.Type, Message: "Сфера компании не задана",
		},
		&validators.StringIsPresent{
			Name: "Company", Field: form.Company, Message: "Название компании не задано",
		},
		&validators.StringIsPresent{
			Name: "Role", Field: form.Role, Message: "Должность не задана",
		},
		&validators.StringIsPresent{
			Name: "Salary", Field: form.Salary, Message: "Зарплата не задана",
		},
	)
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}
	err := h.repository.addVacancy(form)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		component = components.Notification("Произошла ошибка на сервере, попробуйте позднее", components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}
	component = components.Notification("Вакансия успешно создана", components.NotificationSuccess)
	return tadapter.Render(c, component, http.StatusCreated)
}
