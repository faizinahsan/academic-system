package user

import (
	"github.com/faizinahsan/academic-system/internal/controller/http/v1/response"
	"github.com/faizinahsan/academic-system/internal/entity"
	"github.com/gofiber/fiber/v2"
)

// Login handles user login
func (r *User) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := r.validation.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// TODO: Fetch user from DB and verify password
	userEntity := entity.User{
		Username:     req.Username,
		PasswordHash: req.Password, // TODO: Replace with hashed password
	}
	loginRes, err := r.user.Login(c.Context(), userEntity)
	if err != nil {
		return errorResponse(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Login successful", "token": loginRes.Token})
}

// Register handles user registration
func (r *User) RegisterFaker(c *fiber.Ctx) error {
	err := r.user.RegistrationFaker(c.Context())
	if err != nil {
		r.log.Error("Registration error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error{Error: "Internal server error"})
	}
	// TODO: Save user to DB via usecase
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User faker registered successfully"})
}
