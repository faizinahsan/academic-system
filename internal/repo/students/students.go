package students

import (
	"context"
	"fmt"
	"github.com/faizinahsan/academic-system/internal/entity"
	"github.com/faizinahsan/academic-system/pkg/postgres"
)

// New -.
func New(pg *postgres.Postgres) *StudentsRepo {
	return &StudentsRepo{pg}
}

type StudentsRepo struct {
	*postgres.Postgres
}

func (r *StudentsRepo) CreateUserForStudents(ctx context.Context, user *entity.User, student *entity.Students) error {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("StudentsRepo - CreateUserForStudents - Begin: %w", err)
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
		return fmt.Errorf("StudentsRepo - CreateUserForStudents - user ToSql: %w", err)
	}
	_, err = tx.Exec(ctx, userSQL, userArgs...)
	if err != nil {
		return fmt.Errorf("StudentsRepo - CreateUserForStudents - user Exec: %w", err)
	}

	// Insert into students
	studentSQL, studentArgs, err := r.Builder.
		Insert("students").
		Columns("id", "name", "email", "phone", "profile_picture", "gender", "major", "faculty", "created_at", "updated_at").
		Values(student.ID, student.Name, student.Email, student.Phone, student.ProfilePicture, student.Gender, student.Major, student.Faculty, student.CreatedAt, student.UpdatedAt).
		ToSql()
	if err != nil {
		return fmt.Errorf("StudentsRepo - CreateUserForStudents - student ToSql: %w", err)
	}
	_, err = tx.Exec(ctx, studentSQL, studentArgs...)
	if err != nil {
		return fmt.Errorf("StudentsRepo - CreateUserForStudents - student Exec: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("StudentsRepo - CreateUserForStudents - Commit: %w", err)
	}
	return nil
}

func (r *StudentsRepo) GetStudentsList(ctx context.Context) ([]*entity.Students, error) {
	//TODO implement me
	panic("implement me")
}
