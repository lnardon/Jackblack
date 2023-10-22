package RoomManager

import (
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
            client.WriteMessage(websocket.TextMessage,message)
        }
	}
}

func (r *RoomManager)JoinRoom(ws *websocket.Conn, message []byte){
  fmt.Println("JOIN")
  id := ws.RemoteAddr().String()
  r.AllPlayers[id] = ws
  r.Broadcast(ws,message)
}
