package entity

import (
	"image"

	"celestial-odyssey/internal/component"
)

// GameEntity represents a game entity.
type GameEntity struct {
	components map[component.Type]interface{}
}

// Type returns the type of the entity.
func (e *GameEntity) Type() (component.EntityType, bool) {
	t, ok := e.components[component.EntityTypeComponent]
	if !ok {
		return TypeUnknown, false
	}

	return t.(component.EntityType), true
}

// SetType sets the type of the entity.
func (e *GameEntity) SetType(t component.EntityType) {
	e.components[component.EntityTypeComponent] = t
}

// Bounds returns the bounds of the entity.
func (e *GameEntity) Bounds() (image.Rectangle, bool) {
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
func (e *GameEntity) Position() (component.Position, bool) {
	pos, found := e.components[component.PositionComponent]
	if !found {
		return component.Position{}, false
	}

	return pos.(component.Position), true
}

// SetPosition sets the position of the entity.
func (e *GameEntity) SetPosition(position component.Position) {
	e.components[component.PositionComponent] = position
}

// Size returns the size of the entity.
func (e *GameEntity) Size() (component.Size, bool) {
	size, found := e.components[component.SizeComponent]
	if !found {
		return component.Size{}, false
	}

	return size.(component.Size), true
}

// SetSize sets the size of the entity.
func (e *GameEntity) SetSize(size component.Size) {
	e.components[component.SizeComponent] = size
}

// Velocity returns the velocity of the entity.
func (e *GameEntity) Velocity() (component.Velocity, bool) {
	velocity, found := e.components[component.VelocityComponent]
	if !found {
		return component.Velocity{}, false
	}

	return velocity.(component.Velocity), true
}

// SetVelocity sets the velocity of the entity.
func (e *GameEntity) SetVelocity(v component.Velocity) {
	e.components[component.VelocityComponent] = v
}

// Input returns the input state of the entity.
func (e *GameEntity) Input() (component.Input, bool) {
	input, found := e.components[component.InputComponent]
	if !found {
		return component.Input{}, false
	}

	return input.(component.Input), true
}

// SetInput sets the input state of the entity.
func (e *GameEntity) SetInput(input component.Input) {
	e.components[component.InputComponent] = input
}

func newGameEntity() *GameEntity {
	return &GameEntity{
		components: make(map[component.Type]interface{}),
	}
}

func CreatePlayer() *GameEntity {
	player := newGameEntity()

	player.SetType(TypePlayer)
	player.SetPosition(component.Position{X: 0, Y: 0})
	player.SetSize(component.Size{Width: 20, Height: 40})
	player.SetVelocity(component.Velocity{VX: 0, VY: 0})
	player.SetInput(component.Input{Left: false, Right: false, Jump: false})

	return player
}
