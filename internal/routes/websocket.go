package routes

import (
	ws "todo-api/internal/websocket" // Alias your internal package

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func WebSocketRoutes(app *fiber.App) {
	app.Get("/ws", websocket.New(ws.HandleWebSocketConnection))
}
