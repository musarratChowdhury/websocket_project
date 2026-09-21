package main

import "github.com/gorilla/websocket"

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *WebSocketManager
}

func NewClient(conn *websocket.Conn, manager *WebSocketManager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
	}
}
