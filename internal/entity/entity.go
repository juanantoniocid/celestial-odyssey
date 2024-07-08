package entity

import (
	"image"
	"log"

	"celestial-odyssey/internal/component"
)

// GameEntity represents a game entity.
type GameEntity struct {
	components map[component.Kind]interface{}
}

// GetComponent returns the component of the given kind.
func (e *GameEntity) GetComponent(kind component.Kind) interface{} {
	return e.components[kind]
}

// Type returns the type of the entity.
func (e *GameEntity) Type() Type {
	t, ok := e.components[component.TypeComponent].(Type)
	if !ok {
		log.Println("failed to get entity type")
		return TypeUnknown
	}

	return t
}

// Bounds returns the bounds of the entity.
func (e *GameEntity) Bounds() image.Rectangle {
	pos, okPos := e.components[component.PositionComponent].(*component.Position)
	if !okPos {
		log.Println("failed to get box position")
		return image.Rectangle{}
	}

	size, okSize := e.components[component.SizeComponent].(*component.Size)
	if !okSize {
		log.Println("failed to get box size")
		return image.Rectangle{}
	}

	return image.Rect(int(pos.X), int(pos.Y), int(pos.X+size.Width), int(pos.Y+size.Height))
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
		components: make(map[component.Kind]interface{}),
	}
}

func (e *GameEntity) addComponent(kind component.Kind, component interface{}) {
	e.components[kind] = component
}

func CreatePlayer() *GameEntity {
	player := newGameEntity()

	player.addComponent(component.TypeComponent, TypePlayer)
	player.addComponent(component.PositionComponent, &component.Position{X: 0, Y: 0})
	player.addComponent(component.SizeComponent, &component.Size{Width: 20, Height: 40})
	player.addComponent(component.VelocityComponent, &component.Velocity{X: 0, Y: 0})
	player.addComponent(component.InputComponent, &component.Input{Left: false, Right: false, Jump: false})

	return player
}
