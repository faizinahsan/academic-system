package v1

import (
	"github.com/faizinahsan/academic-system/internal/usecase"
	"github.com/faizinahsan/academic-system/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	t usecase.Translation
	l logger.Interface
	v *validator.Validate
}
