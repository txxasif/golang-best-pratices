package routes

import (
	"todo-api/internal/routes/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Create a group for routes with the /api prefix
	api := app.Group("/api")

	// Setup routes within the /api group
	auth.AuthRoutes(api)
	SetupTodoRoutes(api)
	SetupUserRoutes(api)
	WebSocketRoutes(app)

}
