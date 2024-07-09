package entity

import (
	"fmt"
	"image"
	"log"

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
// Returns an error if the type component is not found.
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
func (e *GameEntity) Position() *component.Position {
	pos, ok := e.components[component.PositionComponent].(*component.Position)
	if !ok {
		log.Println("failed to get entity position")
		return &component.Position{}
	}

	return pos
}

// Size returns the size of the entity.
func (e *GameEntity) Size() *component.Size {
	size, ok := e.components[component.SizeComponent].(*component.Size)
	if !ok {
		log.Println("failed to get entity size")
		return &component.Size{}
	}

	return size
}

func newGameEntity() *GameEntity {
	return &GameEntity{
		components: make(map[component.Type]interface{}),
	}
}

func (e *GameEntity) addComponent(kind component.Type, component interface{}) {
	e.components[kind] = component
}

func CreatePlayer() *GameEntity {
	player := newGameEntity()

	player.addComponent(component.EntityTypeComponent, TypePlayer)
	player.addComponent(component.PositionComponent, &component.Position{X: 0, Y: 0})
	player.addComponent(component.SizeComponent, &component.Size{Width: 20, Height: 40})
	player.addComponent(component.VelocityComponent, &component.Velocity{VX: 0, VY: 0})
	player.addComponent(component.InputComponent, &component.Input{Left: false, Right: false, Jump: false})

	return player
}
