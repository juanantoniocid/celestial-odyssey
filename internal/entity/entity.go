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
	Components map[string]interface{}
}

// NewGameEntity creates a new GameEntity instance.
func NewGameEntity(id ID) *GameEntity {
	return &GameEntity{
		ID:         id,
		Components: make(map[string]interface{}),
	}
}

// AddComponent adds a component to the entity.
func (e *GameEntity) AddComponent(name string, component interface{}) {
	e.Components[name] = component
}

// GetComponent returns a component from the entity.
func (e *GameEntity) GetComponent(name string) (interface{}, error) {
	component, exists := e.Components[name]
	if !exists {
		return nil, fmt.Errorf("component %s not found", name)
	}
	return component, nil
}

// RemoveComponent removes a component from the entity.
func (e *GameEntity) RemoveComponent(name string) {
	delete(e.Components, name)
}

func (e *GameEntity) Bounds() image.Rectangle {
	pos, okPos := e.Components["position"].(*component.Position)
	if !okPos {
		log.Println("failed to get box position")
		return image.Rectangle{}
	}

	size, okSize := e.Components["size"].(*component.Size)
	if !okSize {
		log.Println("failed to get box size")
		return image.Rectangle{}
	}

	return image.Rect(int(pos.X), int(pos.Y), int(pos.X+size.Width), int(pos.Y+size.Height))
}
