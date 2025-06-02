package user

import (
	"github.com/faizinahsan/academic-system/internal/usecase"
	"github.com/faizinahsan/academic-system/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// NewTranslationRoutes -.
func NewUserRoutes(apiV1Group fiber.Router,
	user usecase.User,
	log logger.Interface) {
	r := &User{
		user:       user,
		log:        log,
		validation: validator.New(validator.WithRequiredStructEnabled()),
	}

	userGroup := apiV1Group.Group("/user")

	{
		userGroup.Post("/register", r.Register)
		userGroup.Post("/login", r.Login)
	}
}
