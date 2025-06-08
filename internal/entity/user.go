package entity

import (
	faker "github.com/go-faker/faker/v4"
	"time"
)

type User struct {
	Id           int32     `json:"id" faker:"boundary_start=1, boundary_end=1000"`
	Username     string    `json:"username" faker:"username,unique"`
	PasswordHash string    `json:"password_hash" faker:"password"`
	CreatedAt    time.Time `json:"created_at" faker:"-"`
	UpdatedAt    time.Time `json:"updated_at" faker:"-"`
	IsActive     bool      `json:"is_active"`
}

// Example usage
func (user *User) FakeUser() (*User, error) {
	err := faker.FakeData(&user)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.IsActive = true
	return user, err
}

type LoginResponse struct {
	Token string `json:"token"`
}
