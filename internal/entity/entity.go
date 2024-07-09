package entity

import (
	"fmt"
	"image"

	"celestial-odyssey/internal/component"
)

// GameEntity represents a game entity.
type GameEntity struct {
	components map[component.Type]interface{}
}

// Type returns the type of the entity.
func (e *GameEntity) Type() (component.EntityType, error) {
	t, ok := e.components[component.EntityTypeComponent]
	if !ok {
		return TypeUnknown, fmt.Errorf("failed to get entity type")
	}

	return t.(component.EntityType), nil
}

// SetType sets the type of the entity.
func (e *GameEntity) SetType(t component.EntityType) {
	e.components[component.EntityTypeComponent] = t
}

// Bounds returns the bounds of the entity.
func (e *GameEntity) Bounds() (image.Rectangle, error) {
	pos, okPos := e.components[component.PositionComponent]
	if !okPos {
		return image.Rectangle{}, fmt.Errorf("failed to get entity position")
	}

	size, okSize := e.components[component.SizeComponent]
	if !okSize {
		return image.Rectangle{}, fmt.Errorf("failed to get entity size")
	}

	componentPosition := pos.(component.Position)
	componentSize := size.(component.Size)

	rect := image.Rect(
		int(componentPosition.X),
		int(componentPosition.Y),
		int(componentPosition.X+componentSize.Width),
		int(componentPosition.Y+componentSize.Height),
	)
	return rect, nil
}

// Position returns the position of the entity.
func (e *GameEntity) Position() (component.Position, error) {
	pos, ok := e.components[component.PositionComponent]
	if !ok {
		return component.Position{}, fmt.Errorf("failed to get entity position")
	}

	return pos.(component.Position), nil
}

// SetPosition sets the position of the entity.
func (e *GameEntity) SetPosition(position component.Position) {
	e.components[component.PositionComponent] = position
}

// Size returns the size of the entity.
func (e *GameEntity) Size() (component.Size, error) {
	size, ok := e.components[component.SizeComponent]
	if !ok {
		return component.Size{}, fmt.Errorf("failed to get entity size")
	}

	return size.(component.Size), nil
}

// SetSize sets the size of the entity.
func (e *GameEntity) SetSize(size component.Size) {
	e.components[component.SizeComponent] = size
}

// Velocity returns the velocity of the entity.
func (e *GameEntity) Velocity() (component.Velocity, error) {
	velocity, ok := e.components[component.VelocityComponent]
	if !ok {
		return component.Velocity{}, fmt.Errorf("failed to get entity velocity")
	}

	return velocity.(component.Velocity), nil
}

// SetVelocity sets the velocity of the entity.
func (e *GameEntity) SetVelocity(v component.Velocity) {
	e.components[component.VelocityComponent] = v
}

// Input returns the input state of the entity.
func (e *GameEntity) Input() (component.Input, error) {
	input, ok := e.components[component.InputComponent]
	if !ok {
		return component.Input{}, fmt.Errorf("failed to get entity input")
	}

	return input.(component.Input), nil
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
