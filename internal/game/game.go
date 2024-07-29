package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/system/behavior"
	"celestial-odyssey/internal/system/graphics"
	"celestial-odyssey/internal/util"
)

type Game struct {
	levels       []Level
	currentLevel int

	updateSystem behavior.UpdateSystem
	renderer     graphics.Renderer

	dimensions util.Dimensions
}

func NewGame(cfg config.Screen, updateSystem behavior.UpdateSystem, renderer graphics.Renderer) *Game {
	return &Game{
		levels:       make([]Level, 0),
		currentLevel: 0,
		dimensions:   cfg.Dimensions,

		updateSystem: updateSystem,
		renderer:     renderer,
	}
}

func (g *Game) Init() {
	currentLevel := g.levels[g.currentLevel]
	currentLevel.Init()
}

func (g *Game) AddLevel(level Level) {
	g.levels = append(g.levels, level)
}

func (g *Game) Update() error {
	currentLevel := g.levels[g.currentLevel]
	currentSection := currentLevel.CurrentSection()
	entities := currentSection.Entities()

	g.updateSystem.Update(entities)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	currentLevel := g.levels[g.currentLevel]
	currentSection := currentLevel.CurrentSection()
	entities := currentSection.Entities()

	g.renderer.Draw(screen, entities)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.dimensions.Width, g.dimensions.Height
}
