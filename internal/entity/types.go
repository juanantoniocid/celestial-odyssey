package entity

import "celestial-odyssey/internal/component"

// Type represents an entity type.
type Type int

const (
	TypeUnknown component.EntityType = iota
	TypeGround
	TypeBox
	TypePlayer
)
