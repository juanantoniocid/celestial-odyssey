package entities

import (
	"image"

	"celestial-odyssey/util"
)

type Direction int

const (
	Left Direction = iota
	Right
)

type Player struct {
	position      image.Point
	width, height int
	speed         int
	playArea      image.Rectangle
	facing        Direction

	moveLeft, moveRight bool
}

func NewPlayer(dimensions util.Dimensions) *Player {
	return &Player{
		width:  dimensions.Width,
		height: dimensions.Height,
	}
}

func (p *Player) Position() image.Point {
	return p.position
}

func (p *Player) Facing() Direction {
	return p.facing
}

func (p *Player) MoveLeft() {
	p.moveLeft = true
}

func (p *Player) MoveRight() {
	p.moveRight = true
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
		p.facing = Left
	}

	if p.moveRight {
		p.moveRight = false
		p.position.X += p.speed
		p.facing = Right
	}

	p.enforceBoundaries()
}

func (p *Player) enforceBoundaries() {
	if p.position.X < p.playArea.Min.X {
		p.position.X = p.playArea.Min.X
	} else if p.position.X+p.width-1 > p.playArea.Max.X {
		p.position.X = p.playArea.Max.X - p.width + 1
	}
}
