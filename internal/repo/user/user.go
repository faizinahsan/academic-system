package user

import (
	"context"
	"fmt"
	"github.com/faizinahsan/academic-system/internal/entity"
	"github.com/faizinahsan/academic-system/pkg/postgres"
)

const _defaultEntityCap = 64

// TranslationRepo -.
type UserRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

// CreateUser -.
func (r *UserRepo) CreateUser(ctx context.Context, t entity.User) error {
	sql, args, err := r.Builder.
		Insert("users").
		Columns("username, email, phone, password_hash, created_at, updated_at, is_active").
		Values(t.Username, t.Email, t.Phone, t.PasswordHash, t.CreatedAt, t.UpdatedAt, t.IsActive).
		ToSql()
	if err != nil {
		return fmt.Errorf("UserRepo - CreateUser - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserRepo - CreateUser - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *UserRepo) GetUserByID(ctx context.Context, userID string) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}
