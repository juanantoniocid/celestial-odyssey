package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/util"
	"celestial-odyssey/world/entities"
)

type Manager struct {
	currentLevel Level
	renderer     Renderer
	dimensions   util.Dimensions
}

type Renderer interface {
	DrawPlayer(screen *ebiten.Image, player *entities.Player)
}

func NewManager(cfg config.Screen, player *entities.Player, renderer Renderer) *Manager {
	return &Manager{
		currentLevel: NewLevel(cfg, player, renderer),
		renderer:     renderer,
		dimensions:   cfg.Dimensions,
	}
}

func (m *Manager) Update() {
	m.currentLevel.Update()
}

func (m *Manager) Draw(screen *ebiten.Image) {
	m.currentLevel.Draw(screen)
}

func (m *Manager) Layout(_outsideWidth, _outsideHeight int) (screenWidth, screenHeight int) {
	return m.dimensions.Width, m.dimensions.Height
}
