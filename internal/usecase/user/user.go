package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/faizinahsan/academic-system/internal/entity"
	"github.com/faizinahsan/academic-system/internal/repo"
)

// UseCase -.
type UseCase struct {
	repo repo.UserRepo
}

func (u UseCase) UpdatePassword(ctx context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

// New -.
func New(r repo.UserRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (u UseCase) Registration(ctx context.Context, users entity.User) (entity.User, error) {
	if users.Username == "" || users.PasswordHash == "" {
		return entity.User{}, errors.New("all fields are required")
	}
	err := u.repo.CreateUser(ctx, &users)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to create user: %w", err)
	}
	return users, nil

}

func (u UseCase) Login(ctx context.Context, users entity.User) (*entity.LoginResponse, error) {

	userData, err := u.repo.GetUserByID(ctx, users.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if userData.PasswordHash != users.PasswordHash {
		return nil, errors.New("invalid password")
	}
	if !userData.IsActive {
		return nil, errors.New("user is not active")
	}
	token, err := createToken(users.Username)
	if err != nil {
		return nil, errors.New("invalid authentication")
	}
	data := entity.LoginResponse{
		Token: token,
	}
	return &data, nil
}
func (u UseCase) RegistrationFaker(ctx context.Context) error {
	emptyUser := entity.User{}
	fakeUser, _ := emptyUser.FakeUser()
	err := u.repo.CreateUser(ctx, fakeUser)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
func (u UseCase) Profile(ctx context.Context, userID string) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UseCase) Logout(ctx context.Context, token string) error {
	//TODO implement me
	panic("implement me")
}
