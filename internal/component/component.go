package component

// Kind represents the kind of given component.
type Kind int

const (
	// PositionComponent represents a position component.
	PositionComponent Kind = iota
	// SizeComponent represents a size component.
	SizeComponent
	// VelocityComponent represents a velocity component.
	VelocityComponent
	// TypeComponent represents a type component.
	TypeComponent
)

// Position represents a 2D position.
type Position struct {
	X, Y float64
}

// Size represents a 2D size.
type Size struct {
	Width, Height float64
}

// Velocity represents a 2D velocity.
type Velocity struct {
	X, Y float64
}

// Type represents a component type.
type Type int
