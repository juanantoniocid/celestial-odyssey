package entity

type Collection struct {
	entities []*GameEntity
}

// NewCollection creates a new entities' collection.
func NewCollection() *Collection {
	return &Collection{
		entities: make([]*GameEntity, 0),
	}
}

// AddGameEntity adds an entity to the entities' collection.
func (c *Collection) AddGameEntity(entity *GameEntity) {
	c.entities = append(c.entities, entity)
}

// GameEntities returns the entities managed by the entities' collection.
func (c *Collection) GameEntities() []*GameEntity {
	return c.entities
}
