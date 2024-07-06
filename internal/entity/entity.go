package entity

import "fmt"

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
