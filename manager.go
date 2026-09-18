package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type WebSocketManager struct {
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{}
}

func (m *WebSocketManager) serveWS(w http.ResponseWriter, r *http.Request) {
	//upgrade a regular http connection into websocket
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	conn.Close()

}
