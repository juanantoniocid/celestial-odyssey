package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/components"
	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/entities"
	"celestial-odyssey/internal/game"
	"celestial-odyssey/internal/graphics"
	"celestial-odyssey/internal/input"
	"celestial-odyssey/internal/screen"
	"celestial-odyssey/internal/systems"
)

func main() {
	cfg := config.LoadConfig()
	applyWindowSettings(cfg.Window)

	player := createPlayer(cfg.Player)
	renderer := createRenderer(cfg.Player, cfg.Screen, cfg.Ground)
	inputHandler := input.NewKeyboardHandler()
	physicsHandler := systems.NewPhysicsHandler()

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

	box1 := landingSite.CreateEntity()
	box1.AddComponent("position", &components.Position{X: 100, Y: 150})
	box1.AddComponent("size", &components.Size{Width: 100, Height: 50})

	box2 := landingSite.CreateEntity()
	box2.AddComponent("position", &components.Position{X: 120, Y: 50})
	box2.AddComponent("size", &components.Size{Width: 80, Height: 50})

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
