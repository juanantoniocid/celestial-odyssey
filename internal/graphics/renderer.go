package graphics

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/debug"
	"celestial-odyssey/internal/entity"
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

func (r *Renderer) Draw(screen *ebiten.Image, player *entity.Player, character *entity.GameEntity, entities []*entity.GameEntity) {
	r.drawBackground(screen)
	r.drawEntities(screen, entities)
	r.drawPlayer(screen, player)
	r.drawCharacter(screen, character)
}

func (r *Renderer) drawPlayer(screen *ebiten.Image, player *entity.Player) {
	r.op.GeoM.Reset()
	sprite := r.getSprite(player)

	r.op.GeoM.Translate(float64(player.Position().X), float64(player.Position().Y))
	screen.DrawImage(r.playerImages[sprite], r.op)
}

func (r *Renderer) getSprite(player *entity.Player) (spriteType SpriteType) {
	switch player.Action() {
	case entity.ActionIdle:
		spriteType = r.getIdleSprite(player)
	case entity.ActionJumping:
		spriteType = r.getJumpingSprite(player)
	case entity.ActionWalking:
		spriteType = r.getWalkingSprite(player)
	default:
		spriteType = PlayerIdleRight
	}

	return spriteType
}

func (r *Renderer) getIdleSprite(player *entity.Player) SpriteType {
	if player.Direction() == entity.DirectionLeft {
		return PlayerIdleLeft
	}
	return PlayerIdleRight
}

func (r *Renderer) getWalkingSprite(player *entity.Player) SpriteType {
	var frame SpriteType
	frameIndex := player.CurrentStateDuration() / framesPerAnimationFrame % totalWalkingFrames

	switch player.Direction() {
	case entity.DirectionLeft:
		switch frameIndex {
		case 0:
			frame = PlayerWalkingLeft1
		case 1:
			frame = PlayerWalkingLeft2
		case 2:
			frame = PlayerWalkingLeft3
		}
	case entity.DirectionRight:
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

func (r *Renderer) getJumpingSprite(player *entity.Player) SpriteType {
	if player.Direction() == entity.DirectionLeft {
		return PlayerJumpingLeft
	}
	return PlayerJumpingRight
}

func (r *Renderer) drawBackground(screen *ebiten.Image) {
	r.op.GeoM.Reset()
	screen.DrawImage(r.backgroundImage, r.op)
}

func (r *Renderer) drawEntities(screen *ebiten.Image, entities []*entity.GameEntity) {
	for _, e := range entities {
		entityType, err := e.Type()
		if err != nil {
			debug.Log("failed to get entity type: %e", err)
			continue
		}

		switch entityType {
		case entity.TypeBox:
			r.drawBox(screen, e)
		case entity.TypeGround:
			r.drawGround(screen, e)
		default:
			debug.Log("unhandled entity type: %d", entityType)
		}
	}
}

func (r *Renderer) drawBox(screen *ebiten.Image, box *entity.GameEntity) {
	boxBounds, err := box.Bounds()
	if err != nil {
		debug.Log("failed to get box bounds: %e", err)
		return
	}

	r.op.GeoM.Reset()
	r.op.GeoM.Translate(float64(boxBounds.Min.X), float64(boxBounds.Min.Y))

	img := ebiten.NewImage(boxBounds.Dx(), boxBounds.Dy())
	brown := color.RGBA{R: 139, G: 69, B: 19, A: 255}
	img.Fill(brown)

	screen.DrawImage(img, r.op)
}

func (r *Renderer) drawGround(screen *ebiten.Image, ground *entity.GameEntity) {
	groundBounds, err := ground.Bounds()
	if err != nil {
		debug.Log("failed to get ground bounds: %e", err)
		return
	}

	for x := groundBounds.Min.X; x < groundBounds.Dx(); x += r.groundDimensions.Dx() {
		r.op.GeoM.Reset()
		r.op.GeoM.Translate(float64(x), float64(groundBounds.Min.Y))
		screen.DrawImage(r.groundImage, r.op)
	}
}

func (r *Renderer) drawCharacter(screen *ebiten.Image, character *entity.GameEntity) {
	characterPosition, err := character.Position()
	if err != nil {
		debug.Log("failed to get character position: %e", err)
		return
	}

	characterSize, err := character.Size()
	if err != nil {
		debug.Log("failed to get character size: %e", err)
		return
	}

	r.op.GeoM.Reset()
	r.op.GeoM.Translate(characterPosition.X, characterPosition.Y)

	img := ebiten.NewImage(int(characterSize.Width), int(characterSize.Height))
	img.Fill(color.White)

	screen.DrawImage(img, r.op)
}
