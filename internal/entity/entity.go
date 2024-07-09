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

// GetComponent returns the component of the given kind.
func (e *GameEntity) GetComponent(kind component.Type) interface{} {
	return e.components[kind]
}

// Type returns the type of the entity.
func (e *GameEntity) Type() (component.EntityType, error) {
	t, ok := e.components[component.EntityTypeComponent].(component.EntityType)
	if !ok {
		return TypeUnknown, fmt.Errorf("failed to get entity type")
	}
	return t, nil
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

// Size returns the size of the entity.
func (e *GameEntity) Size() (*component.Size, error) {
	size, ok := e.components[component.SizeComponent].(*component.Size)
	if !ok {
		return &component.Size{}, fmt.Errorf("failed to get entity size")
	}

	return size, nil
}

// Velocity returns the velocity of the entity.
func (e *GameEntity) Velocity() (*component.Velocity, error) {
	velocity, ok := e.components[component.VelocityComponent].(*component.Velocity)
	if !ok {
		return &component.Velocity{}, fmt.Errorf("failed to get entity velocity")
	}

	return velocity, nil
}

// Input returns the input state of the entity.
func (e *GameEntity) Input() (*component.Input, error) {
	input, ok := e.components[component.InputComponent].(*component.Input)
	if !ok {
		return &component.Input{}, fmt.Errorf("failed to get entity input")
	}

	return input, nil
}

func newGameEntity() *GameEntity {
	return &GameEntity{
		components: make(map[component.Type]interface{}),
	}
}

func (e *GameEntity) addEntityType(t component.EntityType) {
	e.components[component.EntityTypeComponent] = t
}

func (e *GameEntity) addInput(input *component.Input) {
	e.components[component.InputComponent] = input
}

func (e *GameEntity) addVelocity(v *component.Velocity) {
	e.components[component.VelocityComponent] = v
}

func (e *GameEntity) addSize(s *component.Size) {
	e.components[component.SizeComponent] = s
}

func (e *GameEntity) addPosition(p *component.Position) {
	e.components[component.PositionComponent] = p
}

func CreatePlayer() *GameEntity {
	player := newGameEntity()

	player.addEntityType(TypePlayer)
	player.addPosition(&component.Position{X: 0, Y: 0})
	player.addSize(&component.Size{Width: 20, Height: 40})
	player.addVelocity(&component.Velocity{VX: 0, VY: 0})
	player.addInput(&component.Input{Left: false, Right: false, Jump: false})

	return player
}
