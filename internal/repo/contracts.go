// Package repo implements application outer layer logic. Each logic group in own file.
package repo

import (
	"context"

	"github.com/faizinahsan/academic-system/internal/entity"
)

//go:generate mockgen -source=contracts.go -destination=../usecase/mocks_repo_test.go -package=usecase_test

type (
	// TranslationRepo -.
	TranslationRepo interface {
		Store(context.Context, entity.Translation) error
		GetHistory(context.Context) ([]entity.Translation, error)
	}

	// TranslationWebAPI -.
	TranslationWebAPI interface {
		Translate(entity.Translation) (entity.Translation, error)
	}

	UserRepo interface {
		CreateUser(ctx context.Context, user *entity.User) error
		GetUserByID(ctx context.Context, userID string) (entity.User, error)
	}
	StudentsRepo interface {
		GetStudentsList(ctx context.Context) ([]*entity.Students, error)
		CreateUserForStudents(ctx context.Context, user *entity.User, student *entity.Students) error
	}
	ProfessorsRepo interface {
		GetProfessorsList(ctx context.Context) ([]*entity.Professors, error)
	}
	SBARepo interface {
		GetSBAList(ctx context.Context) ([]*entity.SBA, error)
	}
)
