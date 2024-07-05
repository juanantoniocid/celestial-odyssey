package entities

import "fmt"

type EntityID int

// Entity represents a game entity.
type Entity struct {
	ID         EntityID
	Components map[string]interface{}
}

// NewEntity creates a new Entity instance.
func NewEntity(id EntityID) *Entity {
	return &Entity{
		ID:         id,
		Components: make(map[string]interface{}),
	}
}

// AddComponent adds a component to the entity.
func (e *Entity) AddComponent(name string, component interface{}) {
	e.Components[name] = component
}

// GetComponent returns a component from the entity.
func (e *Entity) GetComponent(name string) (interface{}, error) {
	component, exists := e.Components[name]
	if !exists {
		return nil, fmt.Errorf("component %s not found", name)
	}
	return component, nil
}

// RemoveComponent removes a component from the entity.
func (e *Entity) RemoveComponent(name string) {
	delete(e.Components, name)
}
