package v1

import (
	v1 "github.com/faizinahsan/academic-system/docs/proto/v1"
	"github.com/faizinahsan/academic-system/internal/usecase"
	"github.com/faizinahsan/academic-system/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	v1.TranslationServer

	t usecase.Translation
	l logger.Interface
	v *validator.Validate
}
