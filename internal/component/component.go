package component

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Component represents a component of an entity.
type Component interface{}

// Type represents the kind of given component.
type Type int

const (
	// PositionComponent indicates that the component defines the position of an entity in 2D space.
	PositionComponent Type = iota
	// SizeComponent indicates that the component defines the size (width and height) of an entity.
	SizeComponent
	// VelocityComponent indicates that the component defines the velocity of an entity in 2D space.
	VelocityComponent
	// EntityTypeComponent indicates that the component defines the type of the entity (e.g., player, enemy).
	EntityTypeComponent
	// ActionComponent indicates that the component handles the input state for an entity.
	ActionComponent
	// InputComponent indicates that the component defines a mapping from keys to actions
	InputComponent
	// ColorComponent indicates that the component defines the color of an entity.
	ColorComponent
	// SpriteComponent indicates that the component defines the sprite of an entity.
	SpriteComponent
)

// Position represents a 2D position with X and Y coordinates.
type Position struct {
	X, Y float64
}

// Size represents the dimensions of an entity with width and height.
type Size struct {
	Width, Height float64
}

// Velocity represents the speed and direction of an entity in 2D space with velocities along the X and Y axes.
type Velocity struct {
	VX, VY float64
}

// EntityType represents the type of entity (e.g., player, enemy, box).
type EntityType int

// Action represents the actions that an entity can perform (e.g., move left, move right, jump).s
type Action struct {
	Left, Right, Jump bool
}

// Input represents a mapping from keys to actions.
type Input struct {
	Left, Right, Jump ebiten.Key
}

// Color represents the color of an entity.
type Color struct {
	Color color.RGBA
}

// Sprite represents the sprite of an entity.
type Sprite struct {
	Image *ebiten.Image
}
