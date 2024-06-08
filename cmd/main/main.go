package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/game"
	"celestial-odyssey/internal/graphics"
	"celestial-odyssey/internal/input"
	"celestial-odyssey/internal/physics"
	"celestial-odyssey/internal/screen"
	"celestial-odyssey/internal/util"
	"celestial-odyssey/internal/world/entities"
)

func main() {
	cfg := config.LoadConfig()
	applyWindowSettings(cfg.Window)

	player, playerImages := createPlayer(cfg.Player)
	renderer := createRenderer(cfg.Screen, cfg.Ground, playerImages)
	inputHandler := input.NewKeyboardHandler()
	physicsHandler := physics.NewPhysicsHandler()

	levels := createLevel(cfg.Screen, player, renderer, inputHandler, physicsHandler)
	screenManager := createScreenManager(cfg.Screen, []screen.Level{levels})

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
	player = entities.NewPlayer(cfg)
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

func createRenderer(cfgScreen config.Screen, cfgGround config.Ground, playerImages []*ebiten.Image) *graphics.Renderer {
	renderer := graphics.NewRenderer(playerImages, cfgScreen, cfgGround)

	return renderer
}

func createLevel(cfg config.Screen, player *entities.Player, renderer screen.Renderer, inputHandler screen.InputHandler, physicsHandler screen.PhysicsHandler) screen.Level {
	level := screen.NewLevel()

	landingSite := screen.NewScenario(player, renderer, inputHandler, physicsHandler, cfg.Dimensions.Width, cfg.Dimensions.Height)
	sandDunes := screen.NewScenario(player, renderer, inputHandler, physicsHandler, cfg.Dimensions.Width, cfg.Dimensions.Height)
	ruinedTemple := screen.NewScenario(player, renderer, inputHandler, physicsHandler, cfg.Dimensions.Width, cfg.Dimensions.Height)

	landingSite.AddCollidable(entities.NewBox(image.Rect(100, 100, 200, 200)))

	level.AddScenario(landingSite)
	level.AddScenario(sandDunes)
	level.AddScenario(ruinedTemple)

	return level
}

func createScreenManager(cfg config.Screen, levels []screen.Level) *screen.Manager {
	applyScreenSettings(cfg)
	manager := screen.NewManager(cfg)

	for _, level := range levels {
		manager.AddLevel(level)
	}

	return manager
}

func applyScreenSettings(cfg config.Screen) {
	ebiten.SetScreenClearedEveryFrame(cfg.ClearedEveryFrame)
}

func createGame(screenManager game.ScreenManager) *game.Game {
	g := game.NewGame(screenManager)
	g.Init()

	return g
}
