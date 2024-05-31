package main

import (
	"celestial-odyssey/internal/input"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/game"
	"celestial-odyssey/internal/graphics"
	"celestial-odyssey/internal/screen"
	"celestial-odyssey/internal/util"
	"celestial-odyssey/internal/world/entities"
)

func main() {
	cfg := config.LoadConfig()

	applyWindowSettings(cfg.Window)

	player, playerImages := createPlayer(cfg.Player)

	levels := createLevel(cfg.Screen, player, graphics.NewRenderer(playerImages))

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

func createLevel(cfg config.Screen, player *entities.Player, renderer *graphics.Renderer) screen.Level {
	level := screen.NewLevel()

	landingSiteBg, _, err := ebitenutil.NewImageFromFile("assets/images/scenarios/landing_site.png")
	if err != nil {
		log.Fatal("failed to load landing site background:", err)
	}
	sandDunesBg, _, err := ebitenutil.NewImageFromFile("assets/images/scenarios/sand_dunes.png")
	if err != nil {
		log.Fatal("failed to load sand dunes background:", err)
	}
	ruinedTempleBg, _, err := ebitenutil.NewImageFromFile("assets/images/scenarios/ruined_temple.png")
	if err != nil {
		log.Fatal("failed to load ruined temple background:", err)
	}

	inputHandler := input.NewKeyboardHandler()

	landingSite := screen.NewScenario(player, landingSiteBg, renderer, inputHandler, cfg.Dimensions.Width, cfg.Dimensions.Height)
	sandDunes := screen.NewScenario(player, sandDunesBg, renderer, inputHandler, cfg.Dimensions.Width, cfg.Dimensions.Height)
	ruinedTemple := screen.NewScenario(player, ruinedTempleBg, renderer, inputHandler, cfg.Dimensions.Width, cfg.Dimensions.Height)

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
