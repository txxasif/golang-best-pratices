package handlers

import (
	"strconv"
	"todo-api/internal/models"
	"todo-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

// GetAllUsers fetches all users
func GetAllUsers(c *fiber.Ctx) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}
	return c.JSON(users)
}

// GetUserByID fetches a user by ID
func GetUserByID(c *fiber.Ctx) error {
	// Get the ID from the URL parameter
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // Convert string to uint
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	// Call service to get the user by ID
	user, err := services.GetUserByID(uint(id)) // Pass uint to the service
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	return c.JSON(user)
}

// CreateUser creates a new user
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	createdUser, err := services.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(createdUser)
}
