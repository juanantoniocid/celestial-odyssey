package entity

import "fmt"

type EntityManager struct {
	nextEntityID EntityID
	entities     map[EntityID]*Entity
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		nextEntityID: 0,
		entities:     make(map[EntityID]*Entity),
	}
}

func (em *EntityManager) CreateEntity() *Entity {
	entity := NewEntity(em.nextEntityID)
	em.entities[em.nextEntityID] = entity
	em.nextEntityID++
	return entity
}

func (em *EntityManager) GetEntity(id EntityID) (*Entity, error) {
	entity, exists := em.entities[id]
	if !exists {
		return nil, fmt.Errorf("entity %d not found", id)
	}
	return entity, nil
}

func (em *EntityManager) RemoveEntity(id EntityID) {
	delete(em.entities, id)
}

func (em *EntityManager) Entities() map[EntityID]*Entity {
	return em.entities
}
