package entity

type Entities []*Entity

// NewEntities creates a new entities' collection.
func NewEntities() *Entities {
	entities := make(Entities, 0)
	return &entities
}

// AddEntity adds an entity to the entities' collection.
func (c *Entities) AddEntity(entity *Entity) {
	*c = append(*c, entity)
}

// AddEntities adds entities to the entities' collection.
func (c *Entities) AddEntities(entities *Entities) {
	*c = append(*c, *entities...)
}
