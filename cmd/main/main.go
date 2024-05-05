package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/game"
	"celestial-odyssey/internal/graphics"
	"celestial-odyssey/internal/screen"
	"celestial-odyssey/world/entities"
)

func main() {
	cfg := config.LoadConfig()
	applyWindowSettings(cfg.Window)
	g := createGame(cfg.Game)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func applyWindowSettings(cfg config.Window) {
	ebiten.SetWindowTitle(cfg.Title)
	ebiten.SetWindowSize(cfg.Dimensions.Width, cfg.Dimensions.Height)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetScreenClearedEveryFrame(cfg.ScreenClearedEveryFrame)
}

func createGame(cfg config.Game) *game.Game {
	player, playerImage := createPlayer()
	renderer := graphics.NewRenderer(playerImage)
	screenManager := screen.NewManager(cfg.Dimensions, player, renderer)

	return game.NewGame(cfg.Dimensions, screenManager)
}

func createPlayer() (player *entities.Player, playerImage *ebiten.Image) {
	playerImage = loadPlayerImage()
	player = entities.NewPlayer(playerImage.Bounds().Dx(), playerImage.Bounds().Dy())

	return player, playerImage
}

func loadPlayerImage() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")
	if err != nil {
		log.Fatal(err)
	}

	return img
}
