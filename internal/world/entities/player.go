package entities

import (
	"image"

	"celestial-odyssey/internal/config"
)

const (
	framesPerAnimationFrame = 10
	totalWalkingFrames      = 3
)

type Player struct {
	playArea image.Rectangle
	position image.Point
	width    int
	height   int

	direction   HorizontalDirection
	action      Action
	movingLeft  bool
	movingRight bool
	isJumping   bool

	speed        int
	VelocityY    float64
	jumpVelocity float64
	gravity      float64

	frameIndex   int
	frameCounter int
}

func NewPlayer(cfg config.Player) *Player {
	return &Player{
		playArea: cfg.PlayArea,
		position: image.Point{X: 0, Y: cfg.PlayArea.Max.Y - cfg.Dimensions.Height},
		width:    cfg.Dimensions.Width,
		height:   cfg.Dimensions.Height,

		direction:   Right,
		action:      Idle,
		movingLeft:  false,
		movingRight: false,
		isJumping:   false,

		speed:        cfg.Speed,
		VelocityY:    0,
		jumpVelocity: cfg.JumpVelocity,
		gravity:      cfg.Gravity,

		frameIndex:   0,
		frameCounter: 0,
	}
}

func (p *Player) Bounds() image.Rectangle {
	return image.Rect(p.position.X, p.position.Y, p.position.X+p.width, p.position.Y+p.height)
}

func (p *Player) Position() image.Point {
	return p.position
}

func (p *Player) Width() int {
	return p.width
}

func (p *Player) Height() int {
	return p.height
}

func (p *Player) Direction() HorizontalDirection {
	return p.direction
}

func (p *Player) Action() Action {
	return p.action
}

func (p *Player) FrameIndex() int {
	return p.frameIndex
}

func (p *Player) MoveLeft() {
	p.movingLeft = true
}

func (p *Player) MoveRight() {
	p.movingRight = true
}

func (p *Player) Jump() {
	if !p.isJumping {
		p.isJumping = true
		p.VelocityY = p.jumpVelocity
		p.action = Jumping
	}
}

func (p *Player) Land() {
	p.isJumping = false
	p.VelocityY = 0
	p.action = Idle
}

func (p *Player) Stop() {
	p.movingLeft = false
	p.movingRight = false
	p.action = Idle
}

func (p *Player) SetPositionX(x int) {
	p.position.X = x
}

func (p *Player) SetPositionY(y int) {
	p.position.Y = y
}

func (p *Player) SetPositionAtBottomLeft(pointX int) {
	p.position.X = pointX
}

func (p *Player) SetPositionAtBottomRight(pointX int) {
	p.position.X = pointX - p.width
}

func (p *Player) Update() {
	p.updateHorizontalMovement()
	p.updateVerticalPosition()
	p.updateAnimation()
}

func (p *Player) updateHorizontalMovement() {
	p.action = Idle

	if p.movingLeft {
		p.movingLeft = false
		p.position.X -= p.speed
		p.direction = Left
		p.action = Walking
	} else if p.movingRight {
		p.movingRight = false
		p.position.X += p.speed
		p.direction = Right
		p.action = Walking
	}

	if p.isJumping {
		p.action = Jumping
	}
}

func (p *Player) updateVerticalPosition() {
	if p.isJumping {
		p.VelocityY += p.gravity
		p.position.Y += int(p.VelocityY)

		// Check if player has landed
		if p.position.Y >= p.playArea.Max.Y-p.height {
			p.position.Y = p.playArea.Max.Y - p.height
			p.isJumping = false
			p.VelocityY = 0
		}
	}
}

func (p *Player) updateAnimation() {
	switch p.action {
	case Idle, Jumping:
		p.frameIndex = 0
		p.frameCounter = 0
	case Walking:
		p.updateWalkingAnimation()
	}
}

func (p *Player) updateWalkingAnimation() {
	p.frameCounter++
	if p.frameCounter >= framesPerAnimationFrame {
		p.frameCounter = 0
		p.frameIndex = (p.frameIndex + 1) % totalWalkingFrames
	}
}
