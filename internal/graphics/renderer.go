package graphics

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/world/entities"
)

const (
	framesPerAnimationFrame = 10
	totalWalkingFrames      = 3
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

type RendererImpl struct {
	playerImages []*ebiten.Image
}

func NewRenderer(playerImages []*ebiten.Image) *RendererImpl {
	return &RendererImpl{
		playerImages: playerImages,
	}
}

func (r *RendererImpl) DrawPlayer(screen *ebiten.Image, player *entities.Player) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	var frame SpriteType
	switch player.Action() {
	case entities.ActionIdle:
		frame = r.getIdleSprite(player)
	case entities.ActionJumping:
		frame = r.getJumpingSprite(player)
	case entities.ActionWalking:
		frame = r.getWalkingSprite(player)
	}

	op.GeoM.Translate(float64(player.Position().X), float64(player.Position().Y))
	screen.DrawImage(r.playerImages[frame], op)
}

func (r *RendererImpl) getIdleSprite(player *entities.Player) SpriteType {
	if player.Direction() == entities.DirectionLeft {
		return PlayerIdleLeft
	}
	return PlayerIdleRight
}

func (r *RendererImpl) getWalkingSprite(player *entities.Player) SpriteType {
	var frame SpriteType
	frameIndex := player.CurrentStateDuration() / framesPerAnimationFrame % totalWalkingFrames

	switch player.Direction() {
	case entities.DirectionLeft:
		switch frameIndex {
		case 0:
			frame = PlayerWalkingLeft1
		case 1:
			frame = PlayerWalkingLeft2
		case 2:
			frame = PlayerWalkingLeft3
		}
	case entities.DirectionRight:
		switch frameIndex {
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

func (r *RendererImpl) getJumpingSprite(player *entities.Player) SpriteType {
	if player.Direction() == entities.DirectionLeft {
		return PlayerJumpingLeft
	}
	return PlayerJumpingRight
}

func (r *RendererImpl) DrawBackground(screen *ebiten.Image, background *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	screen.DrawImage(background, op)
}

func (r *RendererImpl) DrawCollidable(screen *ebiten.Image, collidable entities.Collidable) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	bounds := collidable.Bounds()
	op.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))

	img := ebiten.NewImage(bounds.Dx(), bounds.Dy())
	brown := color.RGBA{R: 139, G: 69, B: 19, A: 255}
	img.Fill(brown)

	screen.DrawImage(img, op)
}
