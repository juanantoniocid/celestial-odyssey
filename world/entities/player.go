package entities

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	position    image.Point
	hitbox      image.Rectangle
	levelBounds image.Rectangle
	speed       int

	image *ebiten.Image

	moveLeft  bool
	moveRight bool
}

func NewPlayer(img *ebiten.Image, playArea image.Rectangle) *Player {
	return &Player{
		hitbox:      img.Bounds(),
		image:       img,
		levelBounds: playArea,
	}
}

func (p *Player) MoveLeft() {
	p.moveLeft = true
}

func (p *Player) MoveRight() {
	p.moveRight = true
}

func (p *Player) SetSpeed(speed int) {
	p.speed = speed
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest
	op.GeoM.Translate(float64(p.position.X), float64(p.position.Y))

	screen.DrawImage(p.image, op)
}

func (p *Player) Update() {
	if p.moveLeft {
		p.moveLeft = false
		p.position.X -= p.speed
	}

	if p.moveRight {
		p.moveRight = false
		p.position.X += p.speed
	}

	p.enforceBoundaries()
}

func (p *Player) enforceBoundaries() {
	if p.position.X < p.levelBounds.Min.X {
		p.position.X = p.levelBounds.Min.X
	} else if p.position.X+p.hitbox.Dx() > p.levelBounds.Max.X {
		p.position.X = p.levelBounds.Max.X - p.hitbox.Dx()
	}

	if p.position.Y < p.levelBounds.Min.Y {
		p.position.Y = p.levelBounds.Min.Y
	} else if p.position.Y+p.hitbox.Dy() > p.levelBounds.Max.Y {
		p.position.Y = p.levelBounds.Max.Y - p.hitbox.Dy()
	}
}

func (p *Player) SetGroundRight(ground image.Point) {
	p.position = ground
	p.position.X -= p.hitbox.Dx()
	p.position.Y -= p.hitbox.Dy()
}

func (p *Player) SetGroundLeft(ground image.Point) {
	p.position = ground
	p.position.Y -= p.hitbox.Dy()
}
