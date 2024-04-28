package entities

import (
	"image"
)

type Player struct {
	position      image.Point
	width, height int
	speed         int
	playArea      image.Rectangle

	moveLeft, moveRight bool
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) Position() image.Point {
	return p.position
}

func (p *Player) MoveLeft() {
	p.moveLeft = true
}

func (p *Player) MoveRight() {
	p.moveRight = true
}

func (p *Player) SetDimensions(width, height int) {
	p.width = width
	p.height = height
}

func (p *Player) SetPlayArea(playArea image.Rectangle) {
	p.playArea = playArea
}

func (p *Player) SetSpeed(speed int) {
	p.speed = speed
}

func (p *Player) SetPositionAtBottomLeft(anchorPoint image.Point) {
	p.position = image.Point{X: anchorPoint.X, Y: anchorPoint.Y - p.height}
}

func (p *Player) SetPositionAtBottomRight(anchorPoint image.Point) {
	p.position = image.Point{X: anchorPoint.X - p.width, Y: anchorPoint.Y - p.height}
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
