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

func newGameEntity() *GameEntity {
	return &GameEntity{
		components: make(map[component.Kind]interface{}),
	}
}

func (e *GameEntity) addComponent(kind component.Kind, component interface{}) {
	e.components[kind] = component
}
