package entity

import (
	"image"

	"celestial-odyssey/internal/component"
)

// Entity represents a game entity.
type Entity struct {
	components map[component.Type]component.Component
}

// NewEntity creates a new game entity.
func NewEntity() *Entity {
	return &Entity{
		components: make(map[component.Type]component.Component),
	}
}

// GetComponent returns the component of the entity.
func (e *Entity) GetComponent(t component.Type) (component.Component, bool) {
	c, ok := e.components[t]
	return c, ok
}

// SetComponent sets the component of the entity.
func (e *Entity) SetComponent(t component.Type, c component.Component) {
	e.components[t] = c
}

// RemoveComponent removes the component from the entity.
func (e *Entity) RemoveComponent(t component.Type) {
	delete(e.components, t)
}

// Type returns the type of the entity.
func (e *Entity) Type() (component.EntityType, bool) {
	t, ok := e.components[component.EntityTypeComponent]
	if !ok {
		return TypeUnknown, false
	}

	return t.(component.EntityType), true
}

// SetType sets the type of the entity.
func (e *Entity) SetType(t component.EntityType) {
	e.components[component.EntityTypeComponent] = t
}

// Bounds returns the bounds of the entity.
func (e *Entity) Bounds() (image.Rectangle, bool) {
	pos, found := e.Position()
	if !found {
		return image.Rectangle{}, false
	}

	size, found := e.Size()
	if !found {
		return image.Rectangle{}, false
	}

	rect := image.Rect(
		int(pos.X),
		int(pos.Y),
		int(pos.X+size.Width),
		int(pos.Y+size.Height),
	)
	return rect, true
}

// Position returns the position of the entity.
func (e *Entity) Position() (component.Position, bool) {
	pos, found := e.components[component.PositionComponent]
	if !found {
		return component.Position{}, false
	}

	return pos.(component.Position), true
}

// SetPosition sets the position of the entity.
func (e *Entity) SetPosition(position component.Position) {
	e.components[component.PositionComponent] = position
}

// Size returns the size of the entity.
func (e *Entity) Size() (component.Size, bool) {
	size, found := e.components[component.SizeComponent]
	if !found {
		return component.Size{}, false
	}

	return size.(component.Size), true
}

// SetSize sets the size of the entity.
func (e *Entity) SetSize(size component.Size) {
	e.components[component.SizeComponent] = size
}

// Velocity returns the velocity of the entity.
func (e *Entity) Velocity() (component.Velocity, bool) {
	velocity, found := e.components[component.VelocityComponent]
	if !found {
		return component.Velocity{}, false
	}

	return velocity.(component.Velocity), true
}

// SetVelocity sets the velocity of the entity.
func (e *Entity) SetVelocity(v component.Velocity) {
	e.components[component.VelocityComponent] = v
}

// Action returns the input state of the entity.
func (e *Entity) Action() (component.Action, bool) {
	action, found := e.components[component.ActionComponent]
	if !found {
		return component.Action{}, false
	}

	return action.(component.Action), true
}

// SetAction sets the input state of the entity.
func (e *Entity) SetAction(input component.Action) {
	e.components[component.ActionComponent] = input
}

// Input returns the input map of the entity.
func (e *Entity) Input() (component.Input, bool) {
	input, found := e.components[component.InputComponent]
	if !found {
		return component.Input{}, false
	}

	return input.(component.Input), true
}

// SetInput sets the input map of the entity.
func (e *Entity) SetInput(input component.Input) {
	e.components[component.InputComponent] = input
}

// Color returns the color of the entity.
func (e *Entity) Color() (component.Color, bool) {
	color, found := e.components[component.ColorComponent]
	if !found {
		return component.Color{}, false
	}

	return color.(component.Color), true
}

// SetColor sets the color of the entity.
func (e *Entity) SetColor(color component.Color) {
	e.components[component.ColorComponent] = color
}

// Sprite returns the sprite of the entity.
func (e *Entity) Sprite() (component.Sprite, bool) {
	sprite, found := e.components[component.SpriteComponent]
	if !found {
		return component.Sprite{}, false
	}

	return sprite.(component.Sprite), true
}

// SetSprite sets the sprite of the entity.
func (e *Entity) SetSprite(sprite component.Sprite) {
	e.components[component.SpriteComponent] = sprite
}
