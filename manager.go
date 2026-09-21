package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type WebSocketManager struct {
	sync.RWMutex
	clients ClientList
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients: make(ClientList),
	}
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

// method to add a client to the manager
func (m *WebSocketManager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[client] = true
}

// method to remove a client from the manager
func (m *WebSocketManager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.clients[client]; !ok {
		log.Println("Client not found in manager")
		return
	}
	delete(m.clients, client)
}
