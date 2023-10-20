package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	filePeriod = 10 * time.Second
)

var (
	addr = flag.String("addr", ":8080", "http service address")
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	rooms = make(map[string][]*websocket.Conn)
	mutex sync.RWMutex
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("***** WebSocket Running *****")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	for {
		msgType, p, err := ws.ReadMessage()
		if err != nil {
			return
		}

		if msgType == websocket.TextMessage {
			eventsHandler(string(p), ws)
		}
	}
}

func main() {
	http.HandleFunc("/ws", serveWs)
	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func eventsHandler(event_name string, ws *websocket.Conn) {
	switch event_name {
	case "create_room":
		roomID := fmt.Sprintf("%d", time.Now().UnixNano())
		mutex.Lock()
		rooms[roomID] = []*websocket.Conn{ws}
		mutex.Unlock()

    fmt.Println("--- Room Created ---")

		ws.WriteMessage(websocket.TextMessage, []byte("room_created:"+roomID))
	case "join_room":
		roomID := event_name[len("join_room:"):]
		mutex.Lock()
		if clients, ok := rooms[roomID]; ok {
			rooms[roomID] = append(clients, ws)
		}
		mutex.Unlock()

    fmt.Println("--- Room Joined ---")

		ws.WriteMessage(websocket.TextMessage, []byte("room_joined"))
  // case "get_all_rooms":
  //   all_rooms := []string{}
	// 	mutex.Lock()
  //   for name, _ := range rooms {
  //     all_rooms = append(all_rooms,name)
  //   }
	// 	mutex.Unlock()
    
	// 	ws.WriteMessage(websocket.TextMessage, []byte(all_rooms))
  }
}

func broadcast(roomID string, message []byte) {
	mutex.RLock()
	clients, ok := rooms[roomID]
	mutex.RUnlock()

	if !ok {
		return
	}

	for _, client := range clients {
		client.WriteMessage(websocket.TextMessage, message)
	}
}
