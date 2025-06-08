package students

import (
	"github.com/faizinahsan/academic-system/internal/usecase"
	"github.com/faizinahsan/academic-system/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewStudentsRouter(apiV1Group fiber.Router,
	students usecase.Students,
	l logger.Interface) {
	r := &Students{
		students: students,
		log:      l,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}

	translationGroup := apiV1Group.Group("/students")

	{
		translationGroup.Post("/register", r.Register)
	}
}
