package entity

import (
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

func newGameEntity(id ID) *GameEntity {
	return &GameEntity{
		ID:         id,
		Components: make(map[component.Kind]interface{}),
	}
}

func (e *GameEntity) addComponent(kind component.Kind, component interface{}) {
	e.Components[kind] = component
}

// Bounds returns the bounds of the entity.
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
