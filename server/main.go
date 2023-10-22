package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lnardon/Jackblack/server/RoomManager"
	"github.com/lnardon/Jackblack/server/game"
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
	rooms = make(map[string]*RoomManager.RoomManager)
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

	currentGame := &game.Game{
		HasGameStarted : false,
		CurrentPlayer : "Dealer",
	}

	for {
		msgType, p, err := ws.ReadMessage()
		if err != nil {
		return
		}

		if msgType == websocket.TextMessage {
		eventsHandler(string(p), ws, currentGame)
		}
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("../frontend/dist")))
	http.HandleFunc("/ws", serveWs)
	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func eventsHandler(event_name string, ws *websocket.Conn, currentGame *game.Game) {
	switch strings.Split(event_name, ":")[0] {
		case "create_room":
			roomID := string(event_name[len("create_room:"):])

			fmt.Println(roomID)
			rooms[roomID] = &RoomManager.RoomManager{
				RoomId: roomID,
				RoomSize: "8",
				AllPlayers: make(map[string]*websocket.Conn),
			}
			ws.WriteMessage(websocket.TextMessage, []byte("room_created:"+roomID))
			rooms[roomID].JoinRoom(ws, []byte(ws.RemoteAddr().String()))
			fmt.Println("--- Room Created ---")
		case "join_room":
			roomID := string(event_name[len("join_room:"):])
			currentRoom := rooms[roomID]
			fmt.Println(roomID)
			currentRoom.JoinRoom(ws, []byte(ws.RemoteAddr().String()))
		case "get_all_rooms":
			all_rooms := []string{}
			mutex.Lock()
			for name := range rooms {
				all_rooms = append(all_rooms,name)
			}
			mutex.Unlock()
			var x = []byte{}

			for i:=0; i<len(all_rooms); i++{
				b := []byte(all_rooms[i])
				for j:=0; j<len(b); j++{
					x = append(x,b[j])
				}
			}
			fmt.Println("Room 1: " + all_rooms[0])
			ws.WriteMessage(websocket.TextMessage, x)
		case "start_game":
			fmt.Println("START")
			currentGame.CurrentPlayer = "2"
			currentGame.HasGameStarted = true
		case "handle_game":
			if(currentGame.HasGameStarted){
				fmt.Println("HANDLE GAME")
				if strings.Split(event_name, ":")[2] == "draw_card" {
					roomID := strings.Split(event_name, ":")[1]
					card := currentGame.GetRandomCard()
					rooms[roomID].Broadcast(ws,[]byte(strconv.Itoa(card)))
					 ws.WriteMessage(websocket.TextMessage,[]byte(strconv.Itoa(card)))
				}
			}
  	}
}
