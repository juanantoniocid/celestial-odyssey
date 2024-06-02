package entities

import (
	"image"

	"celestial-odyssey/internal/config"
)

// Player represents the player character in the game.
type Player struct {
	// Position and dimensions
	position image.Point
	width    int
	height   int

	// Movement and actions
	direction     HorizontalDirection
	currentAction CharacterAction
	prevAction    CharacterAction
	isMovingLeft  bool
	isMovingRight bool
	isJumping     bool

	// Physics
	walkingVelocity     int
	verticalVelocity    float64
	initialJumpVelocity float64
	gravity             float64

	// State
	currentStateDuration int
}

// NewPlayer creates a new player character with the given configuration.
func NewPlayer(cfg config.Player) *Player {
	return &Player{
		width:  cfg.Dimensions.Width,
		height: cfg.Dimensions.Height,

		direction:     DirectionRight,
		currentAction: ActionIdle,
		isMovingLeft:  false,
		isMovingRight: false,
		isJumping:     false,

		walkingVelocity:     cfg.WalkingVelocity,
		verticalVelocity:    0,
		initialJumpVelocity: cfg.InitialJumpVelocity,
		gravity:             cfg.Gravity,

		currentStateDuration: 0,
	}
}

// Position and dimensions

// Position returns the current position of the player character.
func (p *Player) Position() image.Point {
	return p.position
}

// SetPositionX sets the x-coordinate of the player character.
func (p *Player) SetPositionX(x int) {
	p.position.X = x
}

// SetPositionY sets the y-coordinate of the player character.
func (p *Player) SetPositionY(y int) {
	p.position.Y = y
}

// Width returns the width of the player character.
func (p *Player) Width() int {
	return p.width
}

// Height returns the height of the player character.
func (p *Player) Height() int {
	return p.height
}

// Bounds returns the bounding box of the player character.
func (p *Player) Bounds() image.Rectangle {
	return image.Rect(p.position.X, p.position.Y, p.position.X+p.width, p.position.Y+p.height)
}

// Movement and actions

// Direction returns the current direction of the player character.
func (p *Player) Direction() HorizontalDirection {
	return p.direction
}

// Action returns the current currentAction of the player character.
func (p *Player) Action() CharacterAction {
	return p.currentAction
}

// MoveLeft moves the player character to the left.
func (p *Player) MoveLeft() {
	p.isMovingLeft = true
}

// MoveRight moves the player character to the right.
func (p *Player) MoveRight() {
	p.isMovingRight = true
}

// Jump makes the player character jump.
func (p *Player) Jump() {
	if !p.isJumping {
		p.isJumping = true
		p.verticalVelocity = p.initialJumpVelocity
		p.currentAction = ActionJumping
	}
}

// IsJumping returns true if the player character is currently jumping.
func (p *Player) IsJumping() bool {
	return p.isJumping
}

// Land makes the player character land on the ground.
func (p *Player) Land() {
	p.isJumping = false
	p.verticalVelocity = 0
	p.currentAction = ActionIdle
}

// Stop stops the player character from moving.
func (p *Player) Stop() {
	p.isMovingLeft = false
	p.isMovingRight = false
	p.currentAction = ActionIdle
}

// Fall makes the player character fall down.
func (p *Player) Fall() {
	p.isJumping = true
	p.verticalVelocity = 0
	p.currentAction = ActionJumping
}

// State

// Update updates the player character's state.
func (p *Player) Update() {
	p.currentAction = ActionIdle
	p.updateHorizontalPosition()
	p.updateVerticalPosition()
	p.updateCurrentStateDuration()
}

// CurrentStateDuration returns the duration of the current state.
func (p *Player) CurrentStateDuration() int {
	return p.currentStateDuration
}

// updateCurrentStateDuration updates the duration of the current state.
func (p *Player) updateCurrentStateDuration() {
	if p.currentAction == p.prevAction {
		p.currentStateDuration++
		return
	}

	p.prevAction = p.currentAction
	p.currentStateDuration = 0
}

// Physics

// updateHorizontalPosition updates the player character's horizontal position.
func (p *Player) updateHorizontalPosition() {
	if p.isMovingLeft {
		p.isMovingLeft = false
		p.position.X -= p.walkingVelocity
		p.direction = DirectionLeft
		p.currentAction = ActionWalking
	} else if p.isMovingRight {
		p.isMovingRight = false
		p.position.X += p.walkingVelocity
		p.direction = DirectionRight
		p.currentAction = ActionWalking
	}
}

// updateVerticalPosition updates the player character's vertical position.
func (p *Player) updateVerticalPosition() {
	if p.isJumping {
		p.currentAction = ActionJumping
		p.verticalVelocity += p.gravity
		p.position.Y += int(p.verticalVelocity)
	}
}
