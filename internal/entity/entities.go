package entity

import (
	"celestial-odyssey/internal/component"
	"celestial-odyssey/internal/config"
)

type Type int

const (
	TypeUnknown component.EntityType = iota
	TypeGround
	TypeBox
	TypePlayer
)

const (
	boxWidth  = 30
	boxHeight = 30

	groundPositionX = 0
	groundPositionY = 172
	groundWidth     = config.ScreenWidth
	groundHeight    = 28
)

type Entities struct {
	entities []*GameEntity
}

// NewEntities creates a new entities' manager.
func NewEntities() *Entities {
	return &Entities{
		entities: make([]*GameEntity, 0),
	}
}

// Entities returns the entities managed by the entities' manager.
func (em *Entities) Entities() []*GameEntity {
	return em.entities
}

// AddBox adds a box entity to the entities' manager.
func (em *Entities) AddBox(x, y float64) {
	box := em.createEntity()

	box.SetType(TypeBox)
	box.SetPosition(component.Position{X: x, Y: y})
	box.SetSize(component.Size{Width: boxWidth, Height: boxHeight})
}

// AddGround adds a ground entity to the entities' manager.
func (em *Entities) AddGround() {
	ground := em.createEntity()

	ground.SetType(TypeGround)
	ground.SetPosition(component.Position{X: groundPositionX, Y: groundPositionY})
	ground.SetSize(component.Size{Width: groundWidth, Height: groundHeight})
}

func (em *Entities) createEntity() *GameEntity {
	entity := newGameEntity()
	em.entities = append(em.entities, entity)

	return entity
}
