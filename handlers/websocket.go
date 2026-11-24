package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed:", err)
		return
	}

	log.Println("WebSocket connection established")
	defer conn.Close()

	log.Println("Reading message from client")
	for {
		messageType, msg, err := conn.ReadMessage()
		if messageType != websocket.TextMessage {
			log.Println("Only text messages are allowed")
			continue
		}
		if err != nil {
			log.Println("Read failed", err)
			break
		}
		log.Printf("Received message: %s", msg)

		err = conn.WriteMessage(
			websocket.TextMessage,
			[]byte("Hello, WebSocket!"),
		)
		if err != nil {
			log.Println("Write failed", err)
			return
		}

		log.Println("Message sent successfully")
	}
}
