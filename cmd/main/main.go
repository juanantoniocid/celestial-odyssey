package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/game"
)

func main() {
	game := game.NewGame()
	ebiten.SetWindowSize(320, 200)
	ebiten.SetWindowTitle("Celestial Odyssey")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
