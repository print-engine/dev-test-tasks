package handlers

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer func(conn *websocket.Conn) {
		time.Sleep(5 * time.Second)
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	err = conn.WriteMessage(websocket.TextMessage, []byte("Connected"))
	if err != nil {
		return
	}
}
