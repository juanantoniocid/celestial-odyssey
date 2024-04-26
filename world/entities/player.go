package entities

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	position image.Point
	bounds   image.Rectangle
	speed    int
	image    *ebiten.Image

	moveLeft  bool
	moveRight bool
}

func NewPlayer(img *ebiten.Image) *Player {
	return &Player{
		bounds: img.Bounds(),
		image:  img,
	}
}

func (p *Player) Top() int {
	return p.position.Y
}

func (p *Player) Bottom() int {
	return p.position.Y + p.bounds.Dy()
}

func (p *Player) Left() int {
	return p.position.X
}

func (p *Player) Right() int {
	return p.position.X + p.bounds.Dx()
}

func (p *Player) MoveLeft() {
	p.moveLeft = true
	p.position.X -= p.speed
}

func (p *Player) MoveRight() {
	p.moveRight = true
	p.position.X += p.speed
}

func (p *Player) MoveToLeftBoundary(origin int) {
	p.position.X = origin
}

func (p *Player) MoveToRightBoundary(width int) {
	p.position.X = width - p.bounds.Dx()
}

func (p *Player) MoveToTopBoundary(origin int) {
	p.position.Y = origin
}

func (p *Player) MoveToBottomBoundary(height int) {
	p.position.Y = height - p.bounds.Dy()
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
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest
	op.GeoM.Translate(float64(p.position.X), float64(p.position.Y))

	screen.DrawImage(p.image, op)
}
