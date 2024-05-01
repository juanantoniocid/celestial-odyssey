package screen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/graphics"
	"celestial-odyssey/world/entities"
)

type Manager struct {
	currentLevel Level
	renderer     *graphics.Renderer
}

func NewScreenManager(width, height int, player *entities.Player, playerImage *ebiten.Image) *Manager {
	renderer := graphics.NewRenderer(playerImage)
	return &Manager{
		currentLevel: NewLevel1(width, height, player, renderer),
		renderer:     renderer,
	}
}

func (m *Manager) Update() {
	m.currentLevel.Update()
}

func (m *Manager) Draw(screen *ebiten.Image) {
	m.currentLevel.Draw(screen)
}
