package graphics

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/util"
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

// Renderer is responsible for drawing the game entities on the screen.
type Renderer struct {
	playerImages     []*ebiten.Image
	backgroundImage  *ebiten.Image
	groundImage      *ebiten.Image
	groundDimensions util.Dimensions

	op *ebiten.DrawImageOptions
}

// NewRenderer creates a new Renderer instance.
func NewRenderer(playerImages []*ebiten.Image, cfgScreen config.Screen, cfgGround config.Ground) *Renderer {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	return &Renderer{
		playerImages:     playerImages,
		backgroundImage:  createBackgroundImage(cfgScreen),
		groundImage:      createGroundImage(cfgGround.File),
		groundDimensions: cfgGround.Dimensions,

		op: op,
	}
}

func createGroundImage(file string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(file)
	if err != nil {
		log.Fatal("failed to load ground image:", err)
		return nil
	}

	return img
}

func createBackgroundImage(cfg config.Screen) *ebiten.Image {
	background := ebiten.NewImage(cfg.Dimensions.Width, cfg.Dimensions.Height)
	background.Fill(cfg.BackgroundColor)

	return background
}

// DrawPlayer draws the player on the screen.
func (r *Renderer) DrawPlayer(screen *ebiten.Image, player *entities.Player) {
	r.op.GeoM.Reset()
	sprite := r.getSprite(player)

	r.op.GeoM.Translate(float64(player.Position().X), float64(player.Position().Y))
	screen.DrawImage(r.playerImages[sprite], r.op)
}

func (r *Renderer) getSprite(player *entities.Player) SpriteType {
	switch player.Action() {
	case entities.ActionIdle:
		return r.getIdleSprite(player)
	case entities.ActionJumping:
		return r.getJumpingSprite(player)
	case entities.ActionWalking:
		return r.getWalkingSprite(player)
	}

	return PlayerIdleRight
}

func (r *Renderer) getIdleSprite(player *entities.Player) SpriteType {
	if player.Direction() == entities.DirectionLeft {
		return PlayerIdleLeft
	}
	return PlayerIdleRight
}

func (r *Renderer) getWalkingSprite(player *entities.Player) SpriteType {
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

func (r *Renderer) getJumpingSprite(player *entities.Player) SpriteType {
	if player.Direction() == entities.DirectionLeft {
		return PlayerJumpingLeft
	}
	return PlayerJumpingRight
}

// DrawBackground draws the background on the screen.
func (r *Renderer) DrawBackground(screen *ebiten.Image, screenWidth, screenHeight int) {
	r.op.GeoM.Reset()
	screen.DrawImage(r.backgroundImage, r.op)

	// Repeat the ground image to fill the screen.
	for x := 0; x < screenWidth; x += r.groundDimensions.Width {
		r.op.GeoM.Reset()
		r.op.GeoM.Translate(float64(x), float64(screenHeight-r.groundDimensions.Height))
		screen.DrawImage(r.groundImage, r.op)
	}
}

// DrawCollidable draws a collidable entity on the screen.
func (r *Renderer) DrawCollidable(screen *ebiten.Image, collidable entities.Collidable) {
	r.op.GeoM.Reset()

	bounds := collidable.Bounds()
	r.op.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))

	img := ebiten.NewImage(bounds.Dx(), bounds.Dy())
	brown := color.RGBA{R: 139, G: 69, B: 19, A: 255}
	img.Fill(brown)

	screen.DrawImage(img, r.op)
}
