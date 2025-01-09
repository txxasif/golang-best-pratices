package websocket

import (
	"fmt"
	"sync"

	"github.com/gofiber/websocket/v2"
)

type Hub struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

var hub *Hub

func InitHub() *Hub {
	if hub == nil {
		hub = &Hub{
			clients: make(map[*websocket.Conn]bool),
		}
	}
	return hub
}

func (h *Hub) RegisterClient(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[conn] = true
	fmt.Println("New client connected!")
}

func (h *Hub) UnregisterClient(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.clients, conn)
	fmt.Println("Client disconnected!")
}

func (h *Hub) BroadcastMessage(message string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for client := range h.clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			client.Close()
			delete(h.clients, client)
		}
	}
}
