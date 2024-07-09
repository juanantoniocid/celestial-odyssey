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
	pos, okPos := e.components[component.PositionComponent].(*component.Position)
	if !okPos {
		return image.Rectangle{}, fmt.Errorf("failed to get entity position")
	}

	size, okSize := e.components[component.SizeComponent].(*component.Size)
	if !okSize {
		return image.Rectangle{}, fmt.Errorf("failed to get entity size")
	}

	rect := image.Rect(int(pos.X), int(pos.Y), int(pos.X+size.Width), int(pos.Y+size.Height))
	return rect, nil
}

// Position returns the position of the entity.
func (e *GameEntity) Position() (*component.Position, error) {
	pos, ok := e.components[component.PositionComponent].(*component.Position)
	if !ok {
		return &component.Position{}, fmt.Errorf("failed to get entity position")
	}

	return pos, nil
}

// SetPosition sets the position of the entity.
func (e *GameEntity) SetPosition(p *component.Position) {
	e.components[component.PositionComponent] = p
}

// Size returns the size of the entity.
func (e *GameEntity) Size() (*component.Size, error) {
	size, ok := e.components[component.SizeComponent].(*component.Size)
	if !ok {
		return &component.Size{}, fmt.Errorf("failed to get entity size")
	}

	return size, nil
}

// SetSize sets the size of the entity.
func (e *GameEntity) SetSize(s *component.Size) {
	e.components[component.SizeComponent] = s
}

// Velocity returns the velocity of the entity.
func (e *GameEntity) Velocity() (*component.Velocity, error) {
	velocity, ok := e.components[component.VelocityComponent].(*component.Velocity)
	if !ok {
		return &component.Velocity{}, fmt.Errorf("failed to get entity velocity")
	}

	return velocity, nil
}

// SetVelocity sets the velocity of the entity.
func (e *GameEntity) SetVelocity(v *component.Velocity) {
	e.components[component.VelocityComponent] = v
}

// Input returns the input state of the entity.
func (e *GameEntity) Input() (*component.Input, error) {
	input, ok := e.components[component.InputComponent].(*component.Input)
	if !ok {
		return &component.Input{}, fmt.Errorf("failed to get entity input")
	}

	return input, nil
}

// SetInput sets the input state of the entity.
func (e *GameEntity) SetInput(input *component.Input) {
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
	player.SetPosition(&component.Position{X: 0, Y: 0})
	player.SetSize(&component.Size{Width: 20, Height: 40})
	player.SetVelocity(&component.Velocity{VX: 0, VY: 0})
	player.SetInput(&component.Input{Left: false, Right: false, Jump: false})

	return player
}
