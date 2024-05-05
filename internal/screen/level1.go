package screen

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/util"
	"celestial-odyssey/world/entities"
)

type Level1 struct {
	player   *entities.Player
	renderer Renderer
}

func NewLevel1(dimensions util.Dimensions, player *entities.Player, renderer Renderer) *Level1 {
	groundLeft := image.Point{X: 0, Y: dimensions.Height}

	player.SetPlayArea(image.Rect(0, 0, dimensions.Width-1, dimensions.Height-1))
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
