package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/factory"
	"celestial-odyssey/internal/game"
	"celestial-odyssey/internal/screen"
	"celestial-odyssey/internal/system"
	"celestial-odyssey/internal/system/graphics"
)

func main() {
	cfg := config.LoadConfig()
	applyWindowSettings(cfg.Window)

	sharedEntities := entity.NewEntities()
	sharedEntities.AddEntity(factory.CreatePlayer())

	inputSystem := system.NewInput()
	actionSystem := system.NewAction()
	gravitySystem := system.NewGravity()
	movementSystem := system.NewMovement()
	systems := system.NewSystems(inputSystem, actionSystem, gravitySystem, movementSystem)

	simpleRenderer := graphics.NewSimpleRenderer()
	rendererManager := graphics.NewRendererManager(simpleRenderer)

	levels := factory.LoadLevel1(sharedEntities, systems, rendererManager)
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

func createScreenManager(cfg config.Screen, levels []screen.Level) *screen.Manager {
	applyScreenSettings(cfg)
	manager := screen.NewManager(cfg)

	for _, l := range levels {
		manager.AddLevel(l)
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
