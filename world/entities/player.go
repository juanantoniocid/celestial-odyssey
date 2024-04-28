package entities

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	position image.Point
	hitbox   image.Rectangle
	playArea image.Rectangle
	speed    int
	image    *ebiten.Image

	moveLeft  bool
	moveRight bool
}

func NewPlayer(img *ebiten.Image, playArea image.Rectangle) *Player {
	return &Player{
		hitbox:   img.Bounds(),
		image:    img,
		playArea: playArea,
	}
}

func (p *Player) MoveLeft() {
	p.moveLeft = true
	p.position.X -= p.speed
}

func (p *Player) MoveRight() {
	p.moveRight = true
	p.position.X += p.speed
}

func (p *Player) SetSpeed(speed int) {
	p.speed = speed
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

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest
	op.GeoM.Translate(float64(p.position.X), float64(p.position.Y))

	screen.DrawImage(p.image, op)
}

func (p *Player) MoveToRightBoundary(width int) {
	p.position.X = width - p.hitbox.Dx()
}

func (p *Player) MoveToBottomBoundary(height int) {
	p.position.Y = height - p.hitbox.Dy()
}

func (p *Player) top() int {
	return p.position.Y
}

func (p *Player) bottom() int {
	return p.position.Y + p.hitbox.Dy()
}

func (p *Player) left() int {
	return p.position.X
}

func (p *Player) right() int {
	return p.position.X + p.hitbox.Dx()
}

func (p *Player) setPositionToLeftBoundary(origin int) {
	p.position.X = origin
}

func (p *Player) setPositionToTopBoundary(origin int) {
	p.position.Y = origin
}

func (p *Player) enforceBoundaries() {
	if p.left() < p.playArea.Min.X {
		p.setPositionToLeftBoundary(p.playArea.Min.X)
	} else if p.right() > p.playArea.Max.X {
		p.MoveToRightBoundary(p.playArea.Max.X)
	}

	if p.top() < p.playArea.Min.Y {
		p.setPositionToTopBoundary(p.playArea.Min.Y)
	} else if p.bottom() > p.playArea.Max.Y {
		p.MoveToBottomBoundary(p.playArea.Max.Y)
	}
}
