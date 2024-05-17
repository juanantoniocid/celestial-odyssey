package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/util"
)

type Level interface {
	Init()
	AddScenario(scenario Scenario)
	Update() error
	Draw(screen *ebiten.Image)
}

type Manager struct {
	levels       []Level
	currentLevel int

	dimensions util.Dimensions
}

func NewManager(cfg config.Screen) *Manager {
	return &Manager{
		levels:       make([]Level, 0),
		currentLevel: 0,
		dimensions:   cfg.Dimensions,
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
	return currentLevel.Update()
}

func (m *Manager) Draw(screen *ebiten.Image) {
	currentLevel := m.levels[m.currentLevel]
	currentLevel.Draw(screen)
}

func (m *Manager) Layout(_outsideWidth, _outsideHeight int) (screenWidth, screenHeight int) {
	return m.dimensions.Width, m.dimensions.Height
}
