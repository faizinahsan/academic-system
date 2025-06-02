package user

import (
	"github.com/faizinahsan/academic-system/internal/entity"
	"github.com/gofiber/fiber/v2"
	"time"
)

// Register handles user registration
func (r *User) Register(c *fiber.Ctx) error {
	type RegisterRequest struct {
		Username string `json:"username" validate:"required,min=3"`
		Email    string `json:"email" validate:"required,email"`
		Phone    string `json:"phone"`
		Password string `json:"password" validate:"required,min=6"`
	}
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := r.validation.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Hash password (use a real hash in production)
	passwordHash := req.Password // TODO: Replace with hash function
	user := entity.User{
		Username:     req.Username,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		IsActive:     true,
	}
	r.log.Info("Registering user: %s", user.Username)
	// TODO: Save user to DB via usecase
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

// Login handles user login
func (r *User) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
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
		Email:        req.Email,
		PasswordHash: req.Password, // TODO: Replace with hashed password
	}
	_, err := r.user.Login(c.Context(), userEntity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})

	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Login successful"})
}
