package services

import (
	"errors"
	"todo-api/internal/models"
)

var tasks []models.Task

// Get all tasks
func GetAllTasks() ([]models.Task, error) {
    return tasks, nil
}

// Create a new task
func CreateTask(task models.Task) (models.Task, error) {
    task.ID = generateID() // Generate a unique ID for the task
    tasks = append(tasks, task)
    return task, nil
}

// Get a task by ID
func GetTaskByID(id uint) (models.Task, error) {
    for _, task := range tasks {
        if task.ID == id {
            return task, nil
        }
    }
    return models.Task{}, errors.New("task not found")
}

// Delete a task by ID
func DeleteTask(id uint) error {
    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            return nil
        }
    }
    return errors.New("task not found")
}

// Helper function to generate task IDs
func generateID() uint {
    if len(tasks) == 0 {
        return 1
    }
    return tasks[len(tasks)-1].ID + 1
}