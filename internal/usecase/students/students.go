package students

import (
	"context"
	"errors"
	"github.com/faizinahsan/academic-system/internal/entity"
	"github.com/faizinahsan/academic-system/internal/repo"
)

// New -.
func New(r repo.StudentsRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

type UseCase struct {
	repo repo.StudentsRepo
}

func (u UseCase) StudentsRegistration(ctx context.Context, students *entity.Students, user *entity.User) (*entity.Students, error) {
	if students.ID == "" || students.Email == "" {
		return nil, errors.New("all fields are required")
	}
	if user.Username == "" || user.PasswordHash == "" {
		return nil, errors.New("all fields are required")
	}
	// Harusnya ada method untuk convert base64 profile picture ke link

	err := u.repo.CreateUserForStudents(ctx, user, students)
	if err != nil {
		return nil, err
	}
	return students, nil
}
