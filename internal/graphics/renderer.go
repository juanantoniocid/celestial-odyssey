package graphics

import (
	entities2 "celestial-odyssey/internal/world/entities"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
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

func (r *Renderer) DrawPlayer(screen *ebiten.Image, player *entities2.Player) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	var frame SpriteType
	switch player.Action() {
	case entities2.ActionIdle:
		frame = r.getIdleSprite(player)
	case entities2.ActionJumping:
		frame = r.getJumpingSprite(player)
	case entities2.ActionWalking:
		frame = r.getWalkingSprite(player)
	}

	op.GeoM.Translate(float64(player.Position().X), float64(player.Position().Y))
	screen.DrawImage(r.playerImages[frame], op)
}

func (r *Renderer) getIdleSprite(player *entities2.Player) SpriteType {
	if player.Direction() == entities2.DirectionLeft {
		return PlayerIdleLeft
	}
	return PlayerIdleRight
}

func (r *Renderer) getWalkingSprite(player *entities2.Player) SpriteType {
	var frame SpriteType

	switch player.Direction() {
	case entities2.DirectionLeft:
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

func (r *Renderer) getJumpingSprite(player *entities2.Player) SpriteType {
	if player.Direction() == entities2.DirectionLeft {
		return PlayerJumpingLeft
	}
	return PlayerJumpingRight
}

func (r *Renderer) DrawBackground(screen *ebiten.Image, background *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	screen.DrawImage(background, op)
}

func (r *Renderer) DrawCollidable(screen *ebiten.Image, collidable entities2.Collidable) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	bounds := collidable.Bounds()
	op.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))

	img := ebiten.NewImage(bounds.Dx(), bounds.Dy())
	brown := color.RGBA{R: 139, G: 69, B: 19, A: 255}
	img.Fill(brown)

	screen.DrawImage(img, op)
}
