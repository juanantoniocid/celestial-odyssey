package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/game"
)

func main() {
	g := game.NewGame()

	screenWidth, screenHeight := 320, 200
	scale := 3.0

	ebiten.SetWindowTitle("Celestial Odyssey")
	ebiten.SetWindowSize(int(float64(screenWidth)*scale), int(float64(screenHeight)*scale))
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetScreenClearedEveryFrame(true)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
