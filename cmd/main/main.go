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
	g := createGame(cfg.Screen, []game.Level{levels}, updateSystem, renderer)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func applyWindowSettings(cfg config.Window) {
	ebiten.SetWindowTitle(cfg.Title)
	ebiten.SetWindowSize(cfg.Dimensions.Width, cfg.Dimensions.Height)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
}

func applyScreenSettings(cfg config.Screen) {
	ebiten.SetScreenClearedEveryFrame(cfg.ClearedEveryFrame)
}

func createGame(cfg config.Screen, levels []game.Level, updateSystem behavior.UpdateSystem, renderer graphics.Renderer) *game.Game {
	applyScreenSettings(cfg)

	g := game.NewGame(cfg, updateSystem, renderer)
	g.Init()

	return g
}
