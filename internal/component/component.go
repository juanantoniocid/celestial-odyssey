package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
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
	// InputComponent indicates that the component handles the input state for an entity.
	InputComponent
	// InputMapComponent indicates that the component defines the mapping from actions to keys.
	InputMapComponent
	// ColorComponent indicates that the component defines the color of an entity.
	ColorComponent
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

// Input represents the input state of an entity, such as movement and action commands.
type Input struct {
	Left, Right, Jump bool
}

// InputMap represents a mapping from actions to keys.
type InputMap struct {
	Left, Right, Jump ebiten.Key
}

// Color represents the color of an entity.
type Color struct {
	Color color.RGBA
}
