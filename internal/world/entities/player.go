package entities

import (
	"image"

	"celestial-odyssey/internal/config"
)

type Player struct {
	playArea image.Rectangle
	position image.Point
	width    int
	height   int

	direction   HorizontalDirection
	action      CharacterAction
	prevAction  CharacterAction
	movingLeft  bool
	movingRight bool
	isJumping   bool

	speed        int
	velocityY    float64
	jumpVelocity float64
	gravity      float64

	stateDuration int
}

func NewPlayer(cfg config.Player) *Player {
	return &Player{
		playArea: cfg.PlayArea,
		position: image.Point{X: 0, Y: cfg.PlayArea.Max.Y - cfg.Dimensions.Height},
		width:    cfg.Dimensions.Width,
		height:   cfg.Dimensions.Height,

		direction:   DirectionRight,
		action:      ActionIdle,
		movingLeft:  false,
		movingRight: false,
		isJumping:   false,

		speed:        cfg.Speed,
		velocityY:    0,
		jumpVelocity: cfg.JumpVelocity,
		gravity:      cfg.Gravity,

		stateDuration: 0,
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

func (p *Player) Action() CharacterAction {
	return p.action
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
		p.velocityY = p.jumpVelocity
		p.action = ActionJumping
	}
}

func (p *Player) Land() {
	p.isJumping = false
	p.velocityY = 0
	p.action = ActionIdle
}

func (p *Player) Stop() {
	p.movingLeft = false
	p.movingRight = false
	p.action = ActionIdle
}

func (p *Player) SetPositionX(x int) {
	p.position.X = x
}

func (p *Player) SetPositionY(y int) {
	p.position.Y = y
}

func (p *Player) StateDuration() int {
	return p.stateDuration
}

func (p *Player) Update() {
	p.updateHorizontalPosition()
	p.updateVerticalPosition()
	p.updateStateDuration()
}

func (p *Player) updateHorizontalPosition() {
	p.action = ActionIdle

	if p.movingLeft {
		p.movingLeft = false
		p.position.X -= p.speed
		p.direction = DirectionLeft
		p.action = ActionWalking
	} else if p.movingRight {
		p.movingRight = false
		p.position.X += p.speed
		p.direction = DirectionRight
		p.action = ActionWalking
	}
}

func (p *Player) updateVerticalPosition() {
	if p.isJumping {
		p.action = ActionJumping
		p.velocityY += p.gravity
		p.position.Y += int(p.velocityY)

		if p.position.Y >= p.playArea.Max.Y-p.height {
			p.position.Y = p.playArea.Max.Y - p.height
			p.isJumping = false
			p.velocityY = 0
		}
	}
}

func (p *Player) updateStateDuration() {
	if p.action == p.prevAction {
		p.stateDuration++
		return
	}

	p.prevAction = p.action
	p.stateDuration = 0
}
