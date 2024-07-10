package entity

import (
	"celestial-odyssey/internal/component"
)

type Type int

const (
	TypeUnknown component.EntityType = iota
	TypeGround
	TypeBox
	TypePlayer
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

// AddEntity adds an entity to the entities' manager.
func (em *Entities) AddEntity(entity *GameEntity) {
	em.entities = append(em.entities, entity)
}

func (em *Entities) createEntity() *GameEntity {
	entity := newGameEntity()
	em.entities = append(em.entities, entity)

	return entity
}
