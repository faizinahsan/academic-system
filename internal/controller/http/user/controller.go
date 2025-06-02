package user

import (
	"github.com/faizinahsan/academic-system/internal/usecase"
	"github.com/faizinahsan/academic-system/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// User -.
type User struct {
	user       usecase.User
	log        logger.Interface
	validation *validator.Validate
}
