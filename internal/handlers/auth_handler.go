package handlers

import (
	"time"
	"todo-api/internal/models"
	"todo-api/internal/services"
	"todo-api/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Check if user already exists
	existingUser, _ := services.GetUserByEmail(user.Email)
	if existingUser.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User already exists"})
	}

	// Generate OTP and set expiry time
	otpCode := utils.GenerateOTP()
	user.OTPCode = otpCode
	user.OTPExpiresAt = time.Now().Add(10 * time.Minute)

	// Save user to DB
	createdUser, err := services.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not register user"})
	}

	// Send OTP via email
	if err := services.SendOTPEmail(createdUser.Email, otpCode); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not send email"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User registered successfully. Check your email for the OTP.",
	})
}
