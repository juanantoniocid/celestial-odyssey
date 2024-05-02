package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"celestial-odyssey/internal/game"
	"celestial-odyssey/internal/graphics"
	"celestial-odyssey/internal/screen"
	"celestial-odyssey/world/entities"
)

const (
	gameTitle    = "Celestial Odyssey"
	windowWidth  = 320
	windowHeight = 200
	scale        = 3.0
)

func main() {
	applyWindowSettings()

	player, playerImage := createPlayer()
	renderer := graphics.NewRenderer(playerImage)
	screenManager := screen.NewManager(windowWidth, windowHeight, player, renderer)

	gameInstance := game.NewGame(windowWidth, windowHeight, screenManager)
	if err := ebiten.RunGame(gameInstance); err != nil {
		log.Fatal(err)
	}
}

func applyWindowSettings() {
	ebiten.SetWindowTitle(gameTitle)
	ebiten.SetWindowSize(int(float64(windowWidth)*scale), int(float64(windowHeight)*scale))
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetScreenClearedEveryFrame(true)
}

func createPlayer() (player *entities.Player, playerImage *ebiten.Image) {
	playerImage = loadPlayerImage()
	player = entities.NewPlayer()
	player.SetDimensions(playerImage.Bounds().Dx(), playerImage.Bounds().Dy())

	return player, playerImage
}

func loadPlayerImage() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")
	if err != nil {
		log.Fatal(err)
	}

	return img
}
