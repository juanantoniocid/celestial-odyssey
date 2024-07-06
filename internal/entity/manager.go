package entity

import "fmt"

type EntityManager struct {
	nextEntityID ID
	entities     map[ID]*GameEntity
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		nextEntityID: 0,
		entities:     make(map[ID]*GameEntity),
	}
}

func (em *EntityManager) CreateEntity() *GameEntity {
	entity := NewGameEntity(em.nextEntityID)
	em.entities[em.nextEntityID] = entity
	em.nextEntityID++
	return entity
}

func (em *EntityManager) GetEntity(id ID) (*GameEntity, error) {
	entity, exists := em.entities[id]
	if !exists {
		return nil, fmt.Errorf("entity %d not found", id)
	}
	return entity, nil
}

func (em *EntityManager) RemoveEntity(id ID) {
	delete(em.entities, id)
}

func (em *EntityManager) Entities() map[ID]*GameEntity {
	return em.entities
}
