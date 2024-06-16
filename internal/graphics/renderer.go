package graphics

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/world/entities"
)

const (
	framesPerAnimationFrame = 10
	imagesInSpriteSheet     = 10
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
	groundDimensions image.Rectangle

	op *ebiten.DrawImageOptions
}

// NewRenderer creates a new Renderer instance.
func NewRenderer(cfgPlayer config.Player, cfgScreen config.Screen, cfgGround config.Ground) *Renderer {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	playerImages := createPlayerImages(cfgPlayer)
	groundImage := createGroundImage(cfgGround.File)

	return &Renderer{
		playerImages:     playerImages,
		backgroundImage:  createBackgroundImage(cfgScreen),
		groundImage:      groundImage,
		groundDimensions: groundImage.Bounds(),

		op: op,
	}
}

func createPlayerImages(cfg config.Player) []*ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(cfg.SpriteSheet)
	if err != nil {
		log.Fatal("failed to load player sprite sheet:", err)
		return nil
	}

	var images []*ebiten.Image
	frameWidth := img.Bounds().Max.X / imagesInSpriteSheet
	frameHeight := img.Bounds().Max.Y
	numFrames := img.Bounds().Max.X / frameWidth

	for i := 0; i < numFrames; i++ {
		x := i * frameWidth
		frame := img.SubImage(image.Rect(x, 0, x+frameWidth, frameHeight)).(*ebiten.Image)
		images = append(images, frame)
	}

	return images
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

func (r *Renderer) Draw(screen *ebiten.Image, world *entities.World) {
	r.drawBackground(screen, 320, 200)
	r.drawBoxes(screen, world.GetBoxes())
	r.drawPlayer(screen, world.GetPlayer())
	r.drawGrounds(screen, world.GetGrounds())
}

func (r *Renderer) drawPlayer(screen *ebiten.Image, player *entities.Player) {
	r.op.GeoM.Reset()
	sprite := r.getSprite(player)

	r.op.GeoM.Translate(float64(player.Position().X), float64(player.Position().Y))
	screen.DrawImage(r.playerImages[sprite], r.op)
}

func (r *Renderer) getSprite(player *entities.Player) (spriteType SpriteType) {
	switch player.Action() {
	case entities.ActionIdle:
		spriteType = r.getIdleSprite(player)
	case entities.ActionJumping:
		spriteType = r.getJumpingSprite(player)
	case entities.ActionWalking:
		spriteType = r.getWalkingSprite(player)
	default:
		spriteType = PlayerIdleRight
	}

	return spriteType
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

func (r *Renderer) drawBackground(screen *ebiten.Image, screenWidth, screenHeight int) {
	r.op.GeoM.Reset()
	screen.DrawImage(r.backgroundImage, r.op)
}

func (r *Renderer) drawGrounds(screen *ebiten.Image, ground []*entities.Ground) {
	for _, g := range ground {
		r.drawGround(screen, g)
	}
}

func (r *Renderer) drawGround(screen *ebiten.Image, ground *entities.Ground) {
	bounds := ground.Bounds()

	for x := bounds.Min.X; x < bounds.Dx(); x += r.groundDimensions.Dx() {
		r.op.GeoM.Reset()
		r.op.GeoM.Translate(float64(x), float64(bounds.Min.Y))
		screen.DrawImage(r.groundImage, r.op)
	}
}

func (r *Renderer) drawBoxes(screen *ebiten.Image, boxes []*entities.Box) {
	for _, b := range boxes {
		r.drawBox(screen, b)
	}
}

func (r *Renderer) drawBox(screen *ebiten.Image, box *entities.Box) {
	r.op.GeoM.Reset()

	bounds := box.Bounds()
	r.op.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))

	img := ebiten.NewImage(bounds.Dx(), bounds.Dy())
	brown := color.RGBA{R: 139, G: 69, B: 19, A: 255}
	img.Fill(brown)

	screen.DrawImage(img, r.op)
}
