// Package v1 implements routing paths. Each services in own file.
package http

import (
	"net/http"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/faizinahsan/academic-system/config"
	_ "github.com/faizinahsan/academic-system/docs" // Swagger docs.
	"github.com/faizinahsan/academic-system/internal/controller/http/middleware"
	studentsRoute "github.com/faizinahsan/academic-system/internal/controller/http/students"
	userRoute "github.com/faizinahsan/academic-system/internal/controller/http/user"
	v1 "github.com/faizinahsan/academic-system/internal/controller/http/v1"
	"github.com/faizinahsan/academic-system/internal/usecase"
	"github.com/faizinahsan/academic-system/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(
	app *fiber.App,
	cfg *config.Config,
	t usecase.Translation,
	l logger.Interface,
	user usecase.User,
	students usecase.Students) {
	// Options
	app.Use(middleware.Logger(l))
	app.Use(middleware.Recovery(l))
	// Prometheus metrics
	if cfg.Metrics.Enabled {
		prometheus := fiberprometheus.New("my-service-name")
		prometheus.RegisterAt(app, "/metrics")
		app.Use(prometheus.Middleware)
	}

	// Swagger
	if cfg.Swagger.Enabled {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	// K8s probe
	app.Get("/healthz", func(ctx *fiber.Ctx) error { return ctx.SendStatus(http.StatusOK) })

	// Routers
	apiV1Group := app.Group("/v1")
	{
		v1.NewTranslationRoutes(apiV1Group, t, l)
	}
	apiUserGroup := app.Group("/v1")
	{
		userRoute.NewUserRoutes(apiUserGroup, user, l)
	}
	apiStudentsGroup := app.Group("/v1")
	{
		studentsRoute.NewStudentsRouter(apiStudentsGroup, students, l)
	}
}
