package students

import (
	"github.com/faizinahsan/academic-system/internal/usecase"
	"github.com/faizinahsan/academic-system/pkg/logger"
	"github.com/go-playground/validator/v10"
)

type Students struct {
	students usecase.Students
	log      logger.Interface
	validate *validator.Validate
}
