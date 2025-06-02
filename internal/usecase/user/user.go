package user

import (
	"context"
	"github.com/faizinahsan/academic-system/internal/entity"
	"github.com/faizinahsan/academic-system/internal/repo"
)

// UseCase -.
type UseCase struct {
	repo repo.UserRepo
}

// New -.
func New(r repo.UserRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (u UseCase) Registration(ctx context.Context, users entity.User) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UseCase) Login(ctx context.Context, users entity.User) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UseCase) Profile(ctx context.Context, userID string) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UseCase) Logout(ctx context.Context, token string) error {
	//TODO implement me
	panic("implement me")
}
