package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/system/behavior"
	"celestial-odyssey/internal/system/graphics"
	"celestial-odyssey/internal/util"
)

// Game represents the game itself.
type Game struct {
	levels       []Level
	currentLevel int

	updateSystem behavior.UpdateSystem
	renderer     graphics.Renderer

	dimensions util.Dimensions
}

// Level represents a game level.
type Level interface {
	// Update updates the level.
	Update()

	// Entities returns the entities in the level.
	Entities() *entity.Entities
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

func (g *Game) Update() error {
	currentLevel := g.levels[g.currentLevel]
	currentEntities := currentLevel.Entities()

	g.updateSystem.Update(currentEntities)
	currentLevel.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	currentLevel := g.levels[g.currentLevel]
	currentEntities := currentLevel.Entities()

	g.renderer.Draw(screen, currentEntities)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.dimensions.Width, g.dimensions.Height
}

func (g *Game) AddLevel(level Level) {
	g.levels = append(g.levels, level)
}
