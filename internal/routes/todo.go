package routes

import (
	"todo-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupTodoRoutes(app *fiber.App) {
	// To-Do Routes
	todo := app.Group("/todos")
	todo.Get("/", handlers.GetAllTasks)
	todo.Post("/", handlers.CreateTask)
	todo.Get("/:id", handlers.GetTaskByID)
	todo.Delete("/:id", handlers.DeleteTask)
}
