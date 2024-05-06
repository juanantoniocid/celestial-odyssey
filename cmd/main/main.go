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

	screenManager := createScreenManager(cfg.Screen)
	g := createGame(screenManager)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func applyWindowSettings(cfg config.Window) {
	ebiten.SetWindowTitle(cfg.Title)
	ebiten.SetWindowSize(cfg.Dimensions.Width, cfg.Dimensions.Height)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
}

func createScreenManager(cfg config.Screen) *screen.Manager {
	applyScreenSettings(cfg)

	player, playerImage := createPlayer()
	renderer := graphics.NewRenderer(playerImage)

	return screen.NewManager(cfg, player, renderer)
}

func applyScreenSettings(cfg config.Screen) {
	ebiten.SetScreenClearedEveryFrame(cfg.ClearedEveryFrame)
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

func createGame(screenManager game.ScreenManager) *game.Game {
	return game.NewGame(screenManager)
}
