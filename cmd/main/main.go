package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/game"
	"celestial-odyssey/internal/graphics"
	"celestial-odyssey/internal/screen"
	"celestial-odyssey/util"
	"celestial-odyssey/world/entities"
)

func main() {
	cfg := config.LoadConfig()

	applyWindowSettings(cfg.Window)

	player, playerImages := createPlayer(cfg.Player)
	screenManager := createScreenManager(cfg.Screen, player, playerImages)
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

func createPlayer(cfg config.Player) (player *entities.Player, playerImages []*ebiten.Image) {
	player = entities.NewPlayer(cfg.Dimensions)
	playerImages = loadPlayerImages(cfg.SpritesFile, cfg.Dimensions)

	return player, playerImages
}

func loadPlayerImages(file string, dimensions util.Dimensions) []*ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(file)
	if err != nil {
		log.Fatal("failed to load player sprite sheet:", err)
		return nil
	}

	var images []*ebiten.Image
	frameWidth := dimensions.Width
	frameHeight := dimensions.Height
	numFrames := img.Bounds().Max.X / frameWidth

	for i := 0; i < numFrames; i++ {
		x := i * frameWidth
		frame := img.SubImage(image.Rect(x, 0, x+frameWidth, frameHeight)).(*ebiten.Image)
		images = append(images, frame)
	}

	return images
}

func createScreenManager(cfg config.Screen, player *entities.Player, playerImages []*ebiten.Image) *screen.Manager {
	applyScreenSettings(cfg)
	renderer := graphics.NewRenderer(playerImages)

	return screen.NewManager(cfg, player, renderer)
}

func applyScreenSettings(cfg config.Screen) {
	ebiten.SetScreenClearedEveryFrame(cfg.ClearedEveryFrame)
}

func createGame(screenManager game.ScreenManager) *game.Game {
	return game.NewGame(screenManager)
}
