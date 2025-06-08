package students

import (
	"github.com/faizinahsan/academic-system/internal/controller/http/students/request"
	"github.com/faizinahsan/academic-system/internal/entity"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

func (r *Students) Register(ctx *fiber.Ctx) error {
	r.log.Info("[%v] request %v called %v", ctx.GetRespHeader(fiber.HeaderXRequestID), "Register", string(ctx.BodyRaw()))

	var body request.StudentRequest

	if err := ctx.BodyParser(&body); err != nil {
		r.log.Error(err, "http - students - Register")

		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	if err := r.validate.Struct(body); err != nil {
		r.log.Error(err, "http - students - Register")

		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	student := &entity.Students{
		ID:             body.NIM,
		Name:           body.Name,
		Email:          body.Email,
		Phone:          body.Phone,
		ProfilePicture: body.ProfilePicture,
		Gender:         body.Gender,
		Major:          body.Major,
		Faculty:        body.Faculty,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	user := &entity.User{
		Username:     body.Username,
		PasswordHash: body.Password,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		IsActive:     false,
	}

	student, err := r.students.StudentsRegistration(
		ctx.UserContext(),
		student,
		user,
	)
	if err != nil {
		r.log.Error(err, "http - students - Register")
		return errorResponse(ctx, http.StatusInternalServerError, "students service problems")
	}
	resStudent := map[string]string{
		"nim": student.ID,
	}

	return ctx.Status(http.StatusOK).JSON(resStudent)
}
