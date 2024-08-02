package game

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/system/behavior"
	"celestial-odyssey/internal/system/graphics"
)

// Game represents the game itself.
type Game struct {
	levels       []Level
	currentLevel int

	updateSystem behavior.UpdateSystem
	renderer     graphics.Renderer

	dimensions config.Dimensions
}

// Level represents a game level.
type Level interface {
	// Update updates the level.
	Update()

	// Entities returns the entities in the level.
	Entities() *entity.Entities
}

// NewGame creates a new game.
func NewGame(cfg config.Screen, updateSystem behavior.UpdateSystem, renderer graphics.Renderer) *Game {
	return &Game{
		levels:       make([]Level, 0),
		currentLevel: 0,
		updateSystem: updateSystem,
		renderer:     renderer,
		dimensions:   cfg.Dimensions,
	}
}

// Update updates the game.
func (g *Game) Update() error {
	if err := g.checkLevelIndex(); err != nil {
		panic(err)
	}

	currentLevel := g.levels[g.currentLevel]
	currentEntities := currentLevel.Entities()

	g.updateSystem.Update(currentEntities)
	currentLevel.Update()

	return nil
}

func (g *Game) checkLevelIndex() error {
	if g.currentLevel < 0 || g.currentLevel >= len(g.levels) {
		return errors.New("current level index is out of range")
	}
	return nil
}

// Draw draws the game.
func (g *Game) Draw(screen *ebiten.Image) {
	if err := g.checkLevelIndex(); err != nil {
		panic(err)
	}

	currentLevel := g.levels[g.currentLevel]
	currentEntities := currentLevel.Entities()

	g.renderer.Draw(screen, currentEntities)
}

// Layout returns the game's dimensions.
func (g *Game) Layout(_outsideWidth, _outsideHeight int) (screenWidth, screenHeight int) {
	return g.dimensions.Width, g.dimensions.Height
}

// AddLevel adds a level to the game.
func (g *Game) AddLevel(level Level) {
	g.levels = append(g.levels, level)
}

func (g *Game) SetCurrentLevel(index int) error {
	if index < 0 || index >= len(g.levels) {
		return errors.New("level index is out of range")
	}

	g.currentLevel = index
	return nil
}
