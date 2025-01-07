package main

import (
	"fmt"
	"log"
	"todo-api/internal/config"
	"todo-api/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables (e.g., DB connection, port)
	cfg := config.Load()

	// Create a new Fiber app
	app := fiber.New()
	// Middleware to log the endpoint being hit

	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("Endpoint hit: %s %s\n", c.Method(), c.Path())
		return c.Next()
	})

	// Register all routes
	routes.SetupRoutes(app)
	host := cfg.Host
	port := cfg.Port
	address := host + ":" + port
	fmt.Println(address)

	// Start the server
	log.Fatal(app.Listen(address))
}
