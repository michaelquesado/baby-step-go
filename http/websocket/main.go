package main

import (
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	Conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		Conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleSocket(ws *websocket.Conn) {
	s.Conns[ws] = true
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		data := buf[:n]
		println(string(data))
		ws.Write([]byte("server responding back"))
	}
}

func main() {
	s := NewServer()
	http.Handle("/connect", websocket.Handler(s.handleSocket))
	http.ListenAndServe(":8080", nil)
}
