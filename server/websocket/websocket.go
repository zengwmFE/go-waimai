package websocket

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

type Message struct {
	Type    int    `json:"type"`
	OrderID uint64 `json:"orderId"`
	Content string `json:"content"`
}

type Hub struct {
	clients map[string]*websocket.Conn
	mu      sync.RWMutex
}

var DefaultHub = &Hub{clients: make(map[string]*websocket.Conn)}

func (h *Hub) HandleConn(c *websocket.Conn, sid string) {
	h.mu.Lock()
	h.clients[sid] = c
	h.mu.Unlock()
}

func (h *Hub) Remove(sid string) {
	h.mu.Lock()
	delete(h.clients, sid)
	h.mu.Unlock()
}

func (h *Hub) Broadcast(msg Message) {
	data, _ := json.Marshal(msg)
	h.mu.RLock()
	defer h.mu.RUnlock()
	for _, conn := range h.clients {
		conn.WriteMessage(websocket.TextMessage, data)
	}
}
