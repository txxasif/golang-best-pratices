package main

import (
	"fmt"
	"log"
	"todo-api/internal/config"
	"todo-api/internal/models"
	"todo-api/internal/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables (e.g., DB connection, port)
	cfg := config.Load()

	// Initialize the database connection
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to the database successfully!")

	// Initialize models and run migrations
	models.Init(db)

	// Create a new Fiber app
	app := fiber.New()

	// Middleware to log the endpoint being hit
	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("Endpoint hit: %s %s\n", c.Method(), c.Path())
		return c.Next()
	})

	// Register all routes
	routes.SetupRoutes(app)

	// // Start the WebSocket hub
	// hub := websocket.InitHub()

	// // Broadcast a message to all connected WebSocket clients every 10 seconds
	// go func() {
	// 	for {
	// 		hub.BroadcastMessage("Hello from the server!")
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()

	// Start the server
	address := cfg.Host + ":" + cfg.Port
	log.Printf("Server is running on %s", address)
	log.Fatal(app.Listen(address))
}
