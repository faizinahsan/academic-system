package seeder

import (
	"github.com/faizinahsan/academic-system/internal/entity"
)

func (s Seed) UserSeed() {
	for i := 0; i < 100; i++ {
		user := &entity.User{}
		user, _ = user.FakeUser()
		_, err := s.db.Exec(
			`INSERT INTO users (username, email, phone, password_hash, created_at, updated_at, is_active) VALUES (?, ?, ?, ?, ?, ?, ?)`,
			user.Username, user.Email, user.Phone, user.PasswordHash, user.CreatedAt, user.UpdatedAt, user.IsActive,
		)
		if err != nil {
			panic(err)
		}
	}
}
