package entity

import (
	"image"

	"celestial-odyssey/internal/component"
)

// Entity represents a game entity.
type Entity struct {
	components map[component.Type]component.Component
}

// NewGameEntity creates a new game entity.
func NewGameEntity() *Entity {
	return &Entity{
		components: make(map[component.Type]component.Component),
	}
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
	pos, found := e.components[component.PositionComponent]
	if !found {
		return image.Rectangle{}, false
	}

	size, found := e.components[component.SizeComponent]
	if !found {
		return image.Rectangle{}, false
	}

	componentPosition := pos.(component.Position)
	componentSize := size.(component.Size)

	rect := image.Rect(
		int(componentPosition.X),
		int(componentPosition.Y),
		int(componentPosition.X+componentSize.Width),
		int(componentPosition.Y+componentSize.Height),
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

// Input returns the input state of the entity.
func (e *Entity) Input() (component.Input, bool) {
	input, found := e.components[component.InputComponent]
	if !found {
		return component.Input{}, false
	}

	return input.(component.Input), true
}

// SetInput sets the input state of the entity.
func (e *Entity) SetInput(input component.Input) {
	e.components[component.InputComponent] = input
}

// InputMap returns the input map of the entity.
func (e *Entity) InputMap() (component.InputMap, bool) {
	inputMap, found := e.components[component.InputMapComponent]
	if !found {
		return component.InputMap{}, false
	}

	return inputMap.(component.InputMap), true
}

// SetInputMap sets the input map of the entity.
func (e *Entity) SetInputMap(inputMap component.InputMap) {
	e.components[component.InputMapComponent] = inputMap
}
