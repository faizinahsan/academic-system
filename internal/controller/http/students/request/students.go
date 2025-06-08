package request

type StudentRequest struct {
	Name            string `json:"name" validate:"required,min=3,max=100" example:"John Doe"`
	Email           string `json:"email" validate:"required,email" example:"example@example.com"`
	Phone           string `json:"phone" validate:"required,min=10,max=15" example:"+1234567890"`
	ProfilePicture  string `json:"profile_picture" validate:"omitempty,url" example:"https://example.com/profile.jpg"`
	NIM             string `json:"nim" validate:"required,min=8,max=12" example:"12345678"`
	Major           string `json:"major" validate:"required,min=3,max=50" example:"Computer Science"`
	Faculty         string `json:"faculty" validate:"required,min=3,max=50" example:"Engineering"`
	Password        string `json:"password" validate:"required,min=6,max=100" example:"securepassword"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password" example:"securepassword"`
	Username        string `json:"username" validate:"required,min=3,max=50" example:"johndoe"`
	Gender          string `json:"gender"`
}
