package RoomManager

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type RoomManager struct {
  RoomId string
  AllPlayers map[string]*websocket.Conn
  RoomSize string
}

func (r *RoomManager)Broadcast(ws *websocket.Conn, message []byte){
  fmt.Println("Broad")
	for _, client := range r.AllPlayers {
        if client.RemoteAddr().String() != ws.RemoteAddr().String(){
            client.WriteMessage(websocket.TextMessage, message)
        }
	}
}

func (r *RoomManager)JoinRoom(ws *websocket.Conn, message []byte){
    fmt.Println("JOIN")
    id := ws.RemoteAddr().String()
    r.AllPlayers[id] = ws
    playerIDs := make([]string, 0, len(r.AllPlayers))
    for playerID := range r.AllPlayers {
        playerIDs = append(playerIDs, playerID)
    }
    jsonData, err := json.Marshal(playerIDs)
    if err != nil {
        fmt.Println("Error encoding player IDs:", err)
        return
    }
    r.Broadcast(ws,jsonData)
    ws.WriteMessage(websocket.TextMessage, jsonData)
}
