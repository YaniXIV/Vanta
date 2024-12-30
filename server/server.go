package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
	"sync"
)

type Server struct {
	mu    sync.Mutex
	conns map[*websocket.Conn]struct{}
	ip    string
}
type IP struct {
	Query string
}

func (i IP) getIp() string {
	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		fmt.Println("Api call failed ", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response ", err)
	}
	json.Unmarshal(body, &i)
	return i.Query
}

func InitServer() *Server {
	fmt.Println("Starting Server!")
	return &Server{
		conns: make(map[*websocket.Conn]struct{}),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("<Server Side> New incoming connection from client:", ws.RemoteAddr())
	s.mu.Lock()
	s.conns[ws] = struct{}{}
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.conns, ws)
		s.mu.Unlock()
	}()
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	//var counter int = 1
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection broke!")
				break
				//connection broke
			}
			fmt.Println("read error", err)
			break
		}
		msg := buf[:n]
		s.echoConn(ws, msg)
		/*
			fmt.Printf("<Server Side> Received message from %v: %s\n", ws.RemoteAddr(), string(msg))
			fmt.Println("we are at read: ", counter)
			counter++
			response := fmt.Sprintf("<Server Side> Message received: %s\n", string(msg))
			if _, err := ws.Write([]byte(response)); err != nil {
				fmt.Println("<Server Side> Write error:", err)
				break
			}

		*/

	}
}
func (s *Server) echoConn(ws *websocket.Conn, m []byte) {
	fmt.Println(s.conns)
	for i := range s.conns {
		if ws == i {
			fmt.Println("ws does == i")
			continue
		}
		if _, err := i.Write(m); err != nil {
			fmt.Println("<Server Side> Write error:", err)
			break
		}
		fmt.Println("ws does not == i")

	}
}

func main() {

	server := InitServer()
	var ip IP
	i := ip.getIp()
	fmt.Printf("Server started at %s", i)
	http.Handle("/ws", websocket.Handler(server.handleWS))
	if err := http.ListenAndServe("0.0.0.0:1444", nil); err != nil {
		fmt.Println("<Server Side> Server error:", err)
	}
}
