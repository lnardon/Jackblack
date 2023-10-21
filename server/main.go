package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
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

func eventsHandler(event_name string, ws *websocket.Conn) {
  currentGame := Game{}
  switch strings.Split(event_name, ":")[0] {
	case "create_room":
		roomID := time.Now().Format("20060102150405")
		mutex.Lock()
		rooms[roomID] = []*websocket.Conn{ws}
		mutex.Unlock()
		ws.WriteMessage(websocket.TextMessage, []byte("room_created:"+roomID))
		
		fmt.Println("--- Room Created ---")
	case "join_room":
		roomID := event_name[len("join_room:"):]
		mutex.Lock()
		if clients, ok := rooms[roomID]; ok {
			rooms[roomID] = append(clients, ws)
			ws.WriteMessage(websocket.TextMessage, []byte("room_joined"))
			fmt.Println("--- Room Joined " + roomID + " ---")
		}
		mutex.Unlock()
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
    card := currentGame.getRandomCard()
    ws.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(card)))
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
    fmt.Println(message)
	}
}

type Game struct {
  currentPlayer string
  allCards [52]string
}

func (g *Game)setCurrentPlayer(id string){
  g.currentPlayer = id 
}

func (g *Game)getRandomCard() (int) {
  randNumber := rand.Intn(52)
  return randNumber
}

func (g *Game)ping(){
  fmt.Println("Flong",g.currentPlayer)
}
