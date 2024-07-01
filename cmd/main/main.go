package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/game"
	"celestial-odyssey/internal/graphics"
	"celestial-odyssey/internal/input"
	"celestial-odyssey/internal/physics"
	"celestial-odyssey/internal/screen"
	"celestial-odyssey/internal/world/entities"
)

func main() {
	cfg := config.LoadConfig()
	applyWindowSettings(cfg.Window)

	player := createPlayer(cfg.Player)
	renderer := createRenderer(cfg.Player, cfg.Screen, cfg.Ground)
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

func createPlayer(cfg config.Player) (player *entities.Player) {
	player = entities.NewPlayer(cfg)

	return player
}

func createRenderer(cfgPlayer config.Player, cfgScreen config.Screen, cfgGround config.Ground) *graphics.Renderer {
	renderer := graphics.NewRenderer(cfgPlayer, cfgScreen, cfgGround)

	return renderer
}

func createLevel(cfg config.Screen, player *entities.Player, renderer screen.Renderer, inputHandler screen.InputHandler, physicsHandler screen.PhysicsHandler) screen.Level {
	level := screen.NewLevel()

	landingSite := screen.NewScenario(player, renderer, inputHandler, physicsHandler, cfg.Dimensions.Width, cfg.Dimensions.Height)
	sandDunes := screen.NewScenario(player, renderer, inputHandler, physicsHandler, cfg.Dimensions.Width, cfg.Dimensions.Height)
	ruinedTemple := screen.NewScenario(player, renderer, inputHandler, physicsHandler, cfg.Dimensions.Width, cfg.Dimensions.Height)

	landingSite.AddBox(entities.NewBox(image.Rect(100, 150, 200, 200)))
	landingSite.AddBox(entities.NewBox(image.Rect(120, 50, 200, 100)))

	landingSite.AddGround(entities.NewGround(image.Rect(0, 172, 320, 200)))

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
