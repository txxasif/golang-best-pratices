package handlers

import (
	"strconv"
	"todo-api/internal/models"
	"todo-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

// Get all tasks
func GetAllTasks(c *fiber.Ctx) error {
    tasks, err := services.GetAllTasks()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch tasks",
        })
    }
    return c.JSON(tasks)
}

// Create a new task
func CreateTask(c *fiber.Ctx) error {
    var task models.Task
    if err := c.BodyParser(&task); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid input",
        })
    }

    createdTask, err := services.CreateTask(task)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create task",
        })
    }
    return c.Status(fiber.StatusCreated).JSON(createdTask)
}

// Get task by ID
func GetTaskByID(c *fiber.Ctx) error {
    id, err := strconv.ParseUint(c.Params("id"), 10, 32)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid ID format",
        })
    }

    task, err := services.GetTaskByID(uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Task not found",
        })
    }
    return c.JSON(task)
}

// Delete task by ID
func DeleteTask(c *fiber.Ctx) error {
    id, err := strconv.ParseUint(c.Params("id"), 10, 32)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid ID format",
        })
    }

    if err := services.DeleteTask(uint(id)); err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Task not found",
        })
    }
    return c.SendStatus(fiber.StatusNoContent)
}