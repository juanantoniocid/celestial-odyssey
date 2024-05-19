package entities

import (
	"image"

	"celestial-odyssey/internal/config"
)

type Direction int

const (
	Left Direction = iota
	Right
)

type Action int

const (
	Idle Action = iota
	Walking
	Jumping
)

type Player struct {
	position      image.Point
	width, height int
	speed         int
	playArea      image.Rectangle

	facing Direction
	action Action

	moveLeft, moveRight, isJumping bool
	frameIndex                     int
	frameCounter                   int

	initialJumpVelocity float64
	velocityY           float64
	gravity             float64
}

func NewPlayer(cfg config.Player) *Player {
	return &Player{
		position: image.Point{X: 0, Y: 0},
		width:    cfg.Dimensions.Width,
		height:   cfg.Dimensions.Height,
		speed:    cfg.Speed,
		playArea: cfg.PlayArea,

		facing: Right,
		action: Idle,

		moveLeft:     false,
		moveRight:    false,
		isJumping:    false,
		frameIndex:   0,
		frameCounter: 0,

		initialJumpVelocity: cfg.InitialJumpVelocity,
		velocityY:           0,
		gravity:             cfg.Gravity,
	}
}

func (p *Player) Position() image.Point {
	return p.position
}

func (p *Player) Width() int {
	return p.width
}

func (p *Player) Facing() Direction {
	return p.facing
}

func (p *Player) Action() Action {
	return p.action
}

func (p *Player) FrameIndex() int {
	return p.frameIndex
}

func (p *Player) MoveLeft() {
	p.moveLeft = true
}

func (p *Player) MoveRight() {
	p.moveRight = true
}

func (p *Player) Jump() {
	if !p.isJumping {
		p.isJumping = true
		p.velocityY = p.initialJumpVelocity
		p.action = Jumping
	}
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
		p.action = Walking
	} else if p.moveRight {
		p.moveRight = false
		p.position.X += p.speed
		p.facing = Right
		p.action = Walking
	} else if p.isJumping {
		p.action = Jumping
	} else {
		p.action = Idle
	}

	p.applyGravity()
	p.updateAnimation()
	p.enforceBoundaries()
}

func (p *Player) applyGravity() {
	if p.isJumping {
		p.velocityY += p.gravity
		p.position.Y += int(p.velocityY)

		// Check if player has landed
		if p.position.Y >= p.playArea.Max.Y-p.height {
			p.position.Y = p.playArea.Max.Y - p.height
			p.isJumping = false
			p.velocityY = 0
		}
	}
}

func (p *Player) updateAnimation() {
	if p.action == Idle {
		p.frameIndex = 0
		return
	}

	framesToUpdate := 10
	framesPerDirection := 3

	p.frameCounter++
	if p.frameCounter >= framesToUpdate {
		p.frameCounter = 0
		p.frameIndex = (p.frameIndex + 1) % framesPerDirection
	}
}

func (p *Player) enforceBoundaries() {
	if p.position.X < p.playArea.Min.X {
		p.position.X = p.playArea.Min.X
	} else if p.position.X+p.width-1 > p.playArea.Max.X {
		p.position.X = p.playArea.Max.X - p.width + 1
	}

	if p.position.Y < p.playArea.Min.Y {
		p.position.Y = p.playArea.Min.Y
	} else if p.position.Y > p.playArea.Max.Y-p.height {
		p.position.Y = p.playArea.Max.Y - p.height
		p.isJumping = false
		p.velocityY = 0
	}
}
