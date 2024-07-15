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
