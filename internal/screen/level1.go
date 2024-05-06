package screen

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/world/entities"
)

type Level1 struct {
	player   *entities.Player
	renderer Renderer
}

func NewLevel1(cfg config.Screen, player *entities.Player, renderer Renderer) *Level1 {
	levelWidth := cfg.Dimensions.Width
	levelHeight := cfg.Dimensions.Height
	groundLeft := image.Point{X: 0, Y: levelHeight - 1}

	player.SetPlayArea(image.Rect(0, 0, levelWidth-1, levelHeight-1))
	player.SetPositionAtBottomLeft(groundLeft)
	player.SetSpeed(2)

	return &Level1{
		player:   player,
		renderer: renderer,
	}
}

func (l1 *Level1) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		l1.player.MoveLeft()
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		l1.player.MoveRight()
	}

	l1.player.Update()
}

func (l1 *Level1) Draw(screen *ebiten.Image) {
	l1.renderer.DrawPlayer(screen, l1.player)
}
