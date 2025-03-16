package vacancy

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"time"
)

type VacancyRepository struct {
	Dbpool       *pgxpool.Pool
	CustomLogger *zerolog.Logger
}

func NewVacancyRepository(dbpool *pgxpool.Pool, customLogger *zerolog.Logger) *VacancyRepository {
	return &VacancyRepository{
		Dbpool:       dbpool,
		CustomLogger: customLogger,
	}
}

func (r *VacancyRepository) addVacancy(form VacancyCreateForm) error {
	query := `INSERT INTO vacancies(email, role, company, salary, type, location, createdat) VALUES (@email, @role, @company, @salary, @type, @location, @createdat)`
	args := pgx.NamedArgs{
		"email":     form.Email,
		"role":      form.Role,
		"company":   form.Company,
		"salary":    form.Salary,
		"type":      form.Type,
		"location":  form.Location,
		"createdat": time.Now(),
	}
	_, err := r.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("невозможно создать вакансию: %v", err)
	}
	return nil
}
