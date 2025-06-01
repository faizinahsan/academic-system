package grpc

import (
	v1 "github.com/faizinahsan/academic-system/internal/controller/grpc/v1"
	"github.com/faizinahsan/academic-system/internal/usecase"
	"github.com/faizinahsan/academic-system/pkg/logger"
	pbgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NewRouter -.
func NewRouter(app *pbgrpc.Server, t usecase.Translation, l logger.Interface) {
	{
		v1.NewTranslationRoutes(app, t, l)
	}

	reflection.Register(app)
}
