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
func (r *UserRepo) CreateUser(ctx context.Context, user *entity.User) error {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("UserRepo - CreateUserWithStudent - Begin: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()
	// Insert into users
	userSQL, userArgs, err := r.Builder.
		Insert("users").
		Columns("username", "password_hash", "created_at", "updated_at", "is_active").
		Values(user.Username, user.PasswordHash, user.CreatedAt, user.UpdatedAt, user.IsActive).
		ToSql()
	if err != nil {
		return fmt.Errorf("UserRepo - CreateUserWithStudent - user ToSql: %w", err)
	}
	_, err = tx.Exec(ctx, userSQL, userArgs...)
	if err != nil {
		return fmt.Errorf("UserRepo - CreateUserWithStudent - user Exec: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("UserRepo - CreateUserWithStudent - Commit: %w", err)
	}
	return nil
}

func (r *UserRepo) GetUserByID(ctx context.Context, userID string) (entity.User, error) {
	sql, args, err := r.Builder.
		Select("users").
		Columns("username, password_hash, created_at, updated_at, is_active").
		Where("username = ?", userID).
		ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - GetUserByID - r.Builder: %w", err)
	}
	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - GetUserByID - r.Pool.Query: %w", err)
	}
	defer rows.Close()
	var user entity.User
	if rows.Next() {
		err = rows.Scan(&user.Username, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt, &user.IsActive)
		if err != nil {
			return entity.User{}, fmt.Errorf("UserRepo - GetUserByID - rows.Scan: %w", err)
		}
	} else {
		return entity.User{}, fmt.Errorf("UserRepo - GetUserByID - user not found")
	}
	return user, nil
}
