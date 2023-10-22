package RoomManager

import (
	"fmt"
	"net"
)

type RoomManager struct {
  RoomId string
  AllPlayers map[net.Addr]string
  RoomSize string
}

func (r *RoomManager)broadcast(){
  fmt.Println("Broad")
}
