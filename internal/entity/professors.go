package entity

type Professors struct {
	ID             string `json:"id" faker:"boundary_start=1, boundary_end=10000000"`
	FirstName      string `json:"first_name" faker:"first_name"`
	LastName       string `json:"last_name" faker:"last_name"`
	Email          string `json:"email" faker:"email"`
	Phone          string `json:"phone" faker:"phone_number"`
	ProfilePicture string `json:"profile_picture" faker:"image_url"`
	Gender         string `json:"gender" faker:"oneof=Male"`
	Major          string `json:"major" faker:"-"`
	Faculty        string `json:"faculty" faker:"-"`
	CreatedAt      string `json:"created_at" faker:"-"`
	UpdatedAt      string `json:"updated_at" faker:"-"`
}
