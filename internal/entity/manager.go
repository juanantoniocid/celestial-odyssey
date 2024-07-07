package entity

import "fmt"

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
