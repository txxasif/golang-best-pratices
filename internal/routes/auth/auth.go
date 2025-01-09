package auth

import (
	"todo-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router) {
	// POST /api/auth/register
	app.Post("/register", handlers.Register)
}
