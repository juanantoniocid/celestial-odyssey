package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/system/behavior"
	"celestial-odyssey/internal/system/graphics"
	"celestial-odyssey/internal/util"
)

type Manager struct {
	levels       []Level
	currentLevel int

	updateSystem behavior.UpdateSystem
	renderer     graphics.Renderer

	dimensions util.Dimensions
}

func NewManager(cfg config.Screen, updateSystem behavior.UpdateSystem, renderer graphics.Renderer) *Manager {
	return &Manager{
		levels:       make([]Level, 0),
		currentLevel: 0,
		dimensions:   cfg.Dimensions,

		updateSystem: updateSystem,
		renderer:     renderer,
	}
}

func (m *Manager) Init() {
	currentLevel := m.levels[m.currentLevel]
	currentLevel.Init()
}

func (m *Manager) AddLevel(level Level) {
	m.levels = append(m.levels, level)
}

func (m *Manager) Update() error {
	currentLevel := m.levels[m.currentLevel]
	currentSection := currentLevel.CurrentSection()
	entities := currentSection.Entities()

	m.updateSystem.Update(entities)
	return nil
}

func (m *Manager) Draw(screen *ebiten.Image) {
	currentLevel := m.levels[m.currentLevel]
	currentSection := currentLevel.CurrentSection()
	entities := currentSection.Entities()

	m.renderer.Draw(screen, entities)
}

func (m *Manager) Layout(_outsideWidth, _outsideHeight int) (screenWidth, screenHeight int) {
	return m.dimensions.Width, m.dimensions.Height
}
