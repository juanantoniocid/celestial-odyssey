package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/factory"
	"celestial-odyssey/internal/game"
	"celestial-odyssey/internal/system/behavior"
	"celestial-odyssey/internal/system/graphics"
)

func main() {
	cfg := config.LoadConfig()
	applyWindowSettings(cfg.Window)

	sharedEntities := entity.NewEntities()
	sharedEntities.AddEntity(factory.CreatePlayer())

	updateSystem := factory.CreateUpdateSystem()
	renderer := factory.CreateRenderer()

	levels := factory.CreateLevel1(sharedEntities)
	screenManager := createScreenManager(cfg.Screen, []game.Level{levels}, updateSystem, renderer)

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

func createScreenManager(cfg config.Screen, levels []game.Level, updateSystem behavior.UpdateSystem, renderer graphics.Renderer) *game.Manager {
	applyScreenSettings(cfg)
	manager := game.NewManager(cfg, updateSystem, renderer)

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
