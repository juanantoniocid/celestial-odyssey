package entity

import (
	"fmt"
	"image"
	"log"

	"celestial-odyssey/internal/component"
)

type ID int

// GameEntity represents a game entity.
type GameEntity struct {
	ID         ID
	Components map[component.Kind]interface{}
}

// NewGameEntity creates a new GameEntity instance.
func NewGameEntity(id ID) *GameEntity {
	return &GameEntity{
		ID:         id,
		Components: make(map[component.Kind]interface{}),
	}
}

// AddComponent adds a component to the entity.
func (e *GameEntity) AddComponent(kind component.Kind, component interface{}) {
	e.Components[kind] = component
}

// GetComponent returns a component from the entity.
func (e *GameEntity) GetComponent(kind component.Kind) (interface{}, error) {
	component, exists := e.Components[kind]
	if !exists {
		return nil, fmt.Errorf("component %d not found", kind)
	}
	return component, nil
}

// RemoveComponent removes a component from the entity.
func (e *GameEntity) RemoveComponent(kind component.Kind) {
	delete(e.Components, kind)
}

func (e *GameEntity) Bounds() image.Rectangle {
	pos, okPos := e.Components[component.PositionComponent].(*component.Position)
	if !okPos {
		log.Println("failed to get box position")
		return image.Rectangle{}
	}

	size, okSize := e.Components[component.SizeComponent].(*component.Size)
	if !okSize {
		log.Println("failed to get box size")
		return image.Rectangle{}
	}

	return image.Rect(int(pos.X), int(pos.Y), int(pos.X+size.Width), int(pos.Y+size.Height))
}
