package main

import "net/http"

func main() {
	startServer(setupAPI())
}

func setupAPI() http.Handler {
	manager := NewWebSocketManager()
	mux := http.NewServeMux()
	mux.Handle("/", loggingMiddleware(http.FileServer(http.Dir("./frontend"))))
	mux.HandleFunc("/ws", manager.serveWS)

	return mux
}
