package main

import (
	"log"

	"celestial-odyssey/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowWidth  = 320
	windowHeight = 200
	scale        = 3.0
)

func main() {
	setWindow()

	gameInstance := game.NewGame(windowWidth, windowHeight)
	if err := ebiten.RunGame(gameInstance); err != nil {
		log.Fatal(err)
	}
}

func setWindow() {
	ebiten.SetWindowTitle("Celestial Odyssey")
	ebiten.SetWindowSize(int(float64(windowWidth)*scale), int(float64(windowHeight)*scale))
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetScreenClearedEveryFrame(true)
}
