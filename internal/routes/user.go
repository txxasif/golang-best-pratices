package routes

import (
	"todo-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app fiber.Router) {
	user := app.Group("/users")
	user.Get("/", handlers.GetAllUsers)
	user.Get("/:id", handlers.GetUserByID)
	user.Post("/", handlers.CreateUser)
}
