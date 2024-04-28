package entities

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	position      image.Point
	width, height int
	speed         int
	playArea      image.Rectangle

	image *ebiten.Image

	moveLeft, moveRight bool
}

func NewPlayer(img *ebiten.Image, playArea image.Rectangle) *Player {
	return &Player{
		width:    img.Bounds().Dx(),
		height:   img.Bounds().Dy(),
		image:    img,
		playArea: playArea,
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
	if p.position.X < p.playArea.Min.X {
		p.position.X = p.playArea.Min.X
	} else if p.position.X+p.width > p.playArea.Max.X {
		p.position.X = p.playArea.Max.X - p.width
	}
}

func (p *Player) SetPositionAtBottomLeft(anchorPoint image.Point) {
	p.position = image.Point{X: anchorPoint.X, Y: anchorPoint.Y - p.height}
}

func (p *Player) SetPositionAtBottomRight(anchorPoint image.Point) {
	p.position = image.Point{X: anchorPoint.X - p.width, Y: anchorPoint.Y - p.height}
}
