package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/world/entities"
)

type SpriteType int

const (
	PlayerWalkingLeft1 SpriteType = iota
	PlayerWalkingLeft2
	PlayerWalkingLeft3
	PlayerIdleLeft
	PlayerIdleRight
	PlayerWalkingRight3
	PlayerWalkingRight2
	PlayerWalkingRight
)

type Renderer struct {
	playerImages []*ebiten.Image
}

func NewRenderer(playerImages []*ebiten.Image) *Renderer {
	return &Renderer{
		playerImages: playerImages,
	}
}

func (r *Renderer) DrawPlayer(screen *ebiten.Image, player *entities.Player) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	frame := PlayerIdleRight
	if player.Facing() == entities.Left {
		frame = PlayerIdleLeft
	}
	op.GeoM.Translate(float64(player.Position().X), float64(player.Position().Y))

	screen.DrawImage(r.playerImages[frame], op)
}
