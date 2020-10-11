package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game"
)

func main() {
	rand.Seed(time.Now().Unix())
	game := shooter.NewGame()
	if err := game.Run(); err != nil {
		log.Fatalf("Game error: %v", err)
	}
}
