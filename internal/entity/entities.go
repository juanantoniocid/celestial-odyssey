package entity

import (
	"celestial-odyssey/internal/component"
	"celestial-odyssey/internal/config"
)

type Type component.Type

const (
	TypeUnknown Type = iota
	TypeGround
	TypeBox
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

func NewEntities() *Entities {
	return &Entities{
		entities: make([]*GameEntity, 0),
	}
}

func (em *Entities) createEntity() *GameEntity {
	entity := newGameEntity()
	return entity
}

func (em *Entities) Entities() []*GameEntity {
	return em.entities
}

func (em *Entities) AddBox(x, y float64) {
	box := em.createEntity()

	box.addComponent(component.TypeComponent, TypeBox)
	box.addComponent(component.PositionComponent, &component.Position{X: x, Y: y})
	box.addComponent(component.SizeComponent, &component.Size{Width: boxWidth, Height: boxHeight})
}

func (em *Entities) AddGround() {
	ground := em.createEntity()

	ground.addComponent(component.TypeComponent, TypeGround)
	ground.addComponent(component.PositionComponent, &component.Position{X: groundPositionX, Y: groundPositionY})
	ground.addComponent(component.SizeComponent, &component.Size{Width: groundWidth, Height: groundHeight})
}
