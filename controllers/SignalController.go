package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	}, // For development only; restrict origin in production
}

var connections = make(map[string][]*websocket.Conn)

// SignalingHandler upgrades HTTP to WebSocket and manages signaling
func SignalingHandler(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("room_id")
	if roomID == "" {
		http.Error(w, "room_id is required", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	// Add connection to the room
	connections[roomID] = append(connections[roomID], conn)

	// Handle messages
	go handleMessages(roomID, conn)
}

func handleMessages(roomID string, conn *websocket.Conn) {
	defer conn.Close()


	for {
		var message map[string]interface{}
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		// Broadcast to all peers in the room
		for _, peer := range connections[roomID] {
			if peer != conn {
				err = peer.WriteJSON(message)
				if err != nil {
					log.Println("Write error:", err)
				}
			}
		}
	}
}
