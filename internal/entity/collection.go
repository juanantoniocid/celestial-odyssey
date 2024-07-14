package entity

type Collection []*Entity

// NewCollection creates a new entities' collection.
func NewCollection() *Collection {
	collection := make(Collection, 0)
	return &collection
}

// AddGameEntity adds an entity to the entities' collection.
func (c *Collection) AddGameEntity(entity *Entity) {
	*c = append(*c, entity)
}
