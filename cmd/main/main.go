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
	images, err := loadPlayerSprites()
	if err != nil {
		log.Fatal("failed to load player sprites:", err)
		return nil
	}

	return images[0]
}

func loadPlayerSprites() ([]*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")
	if err != nil {
		log.Fatal("failed to load player sprite sheet:", err)
		return nil, err
	}

	var sprites []*ebiten.Image
	frameWidth := 16
	frameHeight := 32
	numFrames := img.Bounds().Max.X / frameWidth // Assuming the sheet is a single row

	// Extract each frame from the sprite sheet.
	for i := 0; i < numFrames; i++ {
		x := i * frameWidth
		frame := img.SubImage(image.Rect(x, 0, x+frameWidth, frameHeight)).(*ebiten.Image)
		sprites = append(sprites, frame)
	}

	return sprites, nil
}

func createGame(screenManager game.ScreenManager) *game.Game {
	return game.NewGame(screenManager)
}
