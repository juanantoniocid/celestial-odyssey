package entity

import (
	"fmt"

	"celestial-odyssey/internal/component"
	"celestial-odyssey/internal/config"
)

const (
	groundPositionX = 0
	groundPositionY = 172
	groundWidth     = config.ScreenWidth
	groundHeight    = 28

	boxWidth  = 30
	boxHeight = 30
)

type Entities struct {
	nextEntityID ID
	entities     map[ID]*GameEntity
}

func NewEntities() *Entities {
	return &Entities{
		nextEntityID: 0,
		entities:     make(map[ID]*GameEntity),
	}
}

func (em *Entities) CreateEntity() *GameEntity {
	entity := NewGameEntity(em.nextEntityID)
	em.entities[em.nextEntityID] = entity
	em.nextEntityID++
	return entity
}

func (em *Entities) GetEntity(id ID) (*GameEntity, error) {
	entity, exists := em.entities[id]
	if !exists {
		return nil, fmt.Errorf("entity %d not found", id)
	}
	return entity, nil
}

func (em *Entities) RemoveEntity(id ID) {
	delete(em.entities, id)
}

func (em *Entities) Entities() map[ID]*GameEntity {
	return em.entities
}

func (em *Entities) AddGround() {

	ground := em.CreateEntity()

	ground.AddComponent("type", component.TypeGround)
	ground.AddComponent("position", &component.Position{X: groundPositionX, Y: groundPositionY})
	ground.AddComponent("size", &component.Size{Width: groundWidth, Height: groundHeight})
}

func (em *Entities) AddBox(x, y float64) {
	box := em.CreateEntity()

	box.AddComponent("type", component.TypeBox)
	box.AddComponent("position", &component.Position{X: x, Y: y})
	box.AddComponent("size", &component.Size{Width: boxWidth, Height: boxHeight})
}
