package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	// Create a group for routes with the /api prefix
	api := app.Group("/api")

	// Setup routes within the /api group
	SetupTodoRoutes(api)
	SetupUserRoutes(api)
}
