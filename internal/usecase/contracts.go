// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/faizinahsan/academic-system/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_usecase_test.go -package=usecase_test

type (
	// Translation -.
	Translation interface {
		Translate(context.Context, entity.Translation) (entity.Translation, error)
		History(context.Context) (entity.TranslationHistory, error)
	}

	User interface {
		Registration(ctx context.Context, users entity.User) (entity.User, error)
		Login(ctx context.Context, users entity.User) (*entity.User, error)
		Profile(ctx context.Context, userID string) (entity.User, error)
		Logout(ctx context.Context, token string) error
	}
)
