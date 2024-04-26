package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"celestial-odyssey/internal/game"
	"celestial-odyssey/world/entities"
)

func main() {
	playerImage := loadPlayerImage()
	player := entities.NewPlayer(100, 100, 0, playerImage)

	gameInstance := game.NewGame(player)

	screenWidth, screenHeight := 320, 200
	scale := 3.0

	ebiten.SetWindowTitle("Celestial Odyssey")
	ebiten.SetWindowSize(int(float64(screenWidth)*scale), int(float64(screenHeight)*scale))
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetScreenClearedEveryFrame(true)

	if err := ebiten.RunGame(gameInstance); err != nil {
		log.Fatal(err)
	}
}

func loadPlayerImage() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")
	if err != nil {
		log.Fatal(err)
	}

	return img
}
