package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	To       string `json:"to"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var broadcast = make(chan Message)
var clients = make(map[*websocket.Conn]string)

func main() {

	http.HandleFunc("/ws", handleConnection)
	go handleMessages()

	http.ListenAndServe(":8000", nil)
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	otp := r.URL.Query().Get("otp")
	println(otp)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer ws.Close()
	clients[ws] = otp
	for {
		var msg Message
		if err := ws.ReadJSON(&msg); err != nil {
			delete(clients, ws)
			break
		}

		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		println(msg.Message)
		for client := range clients {
			if clients[client] == msg.To {
				err := client.WriteJSON(msg)
				if err != nil {
					delete(clients, client)
					client.Close()
				}
			}

		}
	}
}
