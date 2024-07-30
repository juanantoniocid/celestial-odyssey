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

// Level represents a game level, which can contain multiple sections.
type Level interface {
	// Init initializes the level with the given parameters.
	Init()

	// AddSection adds a section to the level.
	AddSection(section Section)

	// CurrentSection returns the current section.
	CurrentSection() Section
}

// Section represents a single scenario within a level.
type Section interface {
	// Entities returns the entities in the scenario.
	Entities() *entity.Entities

	// SetPlayerPositionAtLeft sets the player position on the left side of the screen.
	SetPlayerPositionAtLeft()
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
	g.updateSystem.Update(g.sectionEntities())
	return nil
}

func (g *Game) sectionEntities() *entity.Entities {
	currentLevel := g.levels[g.currentLevel]
	currentSection := currentLevel.CurrentSection()
	return currentSection.Entities()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.Draw(screen, g.sectionEntities())
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.dimensions.Width, g.dimensions.Height
}
