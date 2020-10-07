package main

import (
	"log"

	"github.com/develersrl/golab2020-go-game-dev/shooter-game"
)

func main() {
	game := shooter.NewGame()
	if err := game.Run(); err != nil {
		log.Fatalf("Game error: %v", err)
	}
}
