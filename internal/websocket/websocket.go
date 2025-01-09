package websocket

import (
	"fmt"

	"github.com/gofiber/websocket/v2"
)

func HandleWebSocketConnection(c *websocket.Conn) {
	hub := InitHub()
	hub.RegisterClient(c)
	defer hub.UnregisterClient(c)

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}
		fmt.Printf("Received: %s\n", msg)
		hub.BroadcastMessage(string(msg))
	}
}
