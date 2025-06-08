package entity

import "time"

type Students struct {
	ID             string    `json:"id" faker:"boundary_start=1, boundary_end=10000000"`
	Name           string    `json:"name" faker:"name"`
	Email          string    `json:"email" faker:"email"`
	Phone          string    `json:"phone" faker:"phone_number"`
	ProfilePicture string    `json:"profile_picture" faker:"image_url"`
	Gender         string    `json:"gender" faker:"oneof=Male"`
	Major          string    `json:"major" faker:"-"`
	Faculty        string    `json:"faculty" faker:"-"`
	CreatedAt      time.Time `json:"created_at" faker:"datetime"`
	UpdatedAt      time.Time `json:"updated_at" faker:"datetime"`
}
