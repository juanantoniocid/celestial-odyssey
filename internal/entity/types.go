package entity

import "image"

type HorizontalDirection int

const (
	DirectionLeft HorizontalDirection = iota
	DirectionRight
)

type CharacterAction int

const (
	ActionIdle CharacterAction = iota
	ActionWalking
	ActionJumping
)

// Collidable is an interface that represents an entity that can collide with other entities.
type Collidable interface {
	Bounds() image.Rectangle
}
