package game

import (
	"fmt"
	"math/rand"
)

type Game struct {
  CurrentPlayer string
  allCards [52]string
  HasGameStarted bool
  
}

func (g *Game)setCurrentPlayer(id string){
  g.CurrentPlayer = id 
}

func (g *Game)GetRandomCard() (int) {
  randNumber := rand.Intn(52)
  return randNumber
}

func (g *Game)drawCard()int{
	fmt.Println("Flong",g.CurrentPlayer)
	card := g.GetRandomCard()
	return card
}

func (g *Game)ping(){
  fmt.Println("Flong",g.CurrentPlayer)
}
