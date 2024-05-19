package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/world/entities"
)

type SpriteType int

const (
	PlayerJumpingLeft SpriteType = iota
	PlayerWalkingLeft1
	PlayerWalkingLeft2
	PlayerWalkingLeft3
	PlayerIdleLeft
	PlayerIdleRight
	PlayerWalkingRight3
	PlayerWalkingRight2
	PlayerWalkingRight1
	PlayerJumpingRight
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

	var frame SpriteType
	switch player.Action() {
	case entities.Idle:
		frame = r.getIdleSprite(player)
	case entities.Jumping:
		frame = r.getJumpingSprite(player)
	case entities.Walking:
		frame = r.getWalkingSprite(player)
	}

	op.GeoM.Translate(float64(player.Position().X), float64(player.Position().Y))
	screen.DrawImage(r.playerImages[frame], op)
}

func (r *Renderer) getIdleSprite(player *entities.Player) SpriteType {
	if player.Direction() == entities.Left {
		return PlayerIdleLeft
	}
	return PlayerIdleRight
}

func (r *Renderer) getWalkingSprite(player *entities.Player) SpriteType {
	var frame SpriteType

	switch player.Direction() {
	case entities.Left:
		switch player.FrameIndex() {
		case 0:
			frame = PlayerWalkingLeft1
		case 1:
			frame = PlayerWalkingLeft2
		case 2:
			frame = PlayerWalkingLeft3
		}
	default:
		switch player.FrameIndex() {
		case 0:
			frame = PlayerWalkingRight1
		case 1:
			frame = PlayerWalkingRight2
		case 2:
			frame = PlayerWalkingRight3
		}
	}

	return frame
}

func (r *Renderer) getJumpingSprite(player *entities.Player) SpriteType {
	if player.Direction() == entities.Left {
		return PlayerJumpingLeft
	}
	return PlayerJumpingRight
}

func (r *Renderer) DrawBackground(screen *ebiten.Image, background *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	screen.DrawImage(background, op)
}
