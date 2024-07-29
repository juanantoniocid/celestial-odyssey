package factory

import (
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/game"
)

func CreateLevel1(sharedEntities *entity.Entities) game.Level {
	level1 := game.NewLevel()

	level1.AddSection(CreateLevel1Section1(sharedEntities))
	level1.AddSection(CreateLevel1Section2(sharedEntities))

	return level1
}

func CreateLevel1Section1(sharedEntities *entity.Entities) game.Section {
	entities := entity.NewEntities()
	entities.AddEntities(sharedEntities)

	ground := CreateGround()
	entities.AddEntity(ground)

	box1 := CreateBox(100, 130)
	entities.AddEntity(box1)

	box2 := CreateBox(120, 50)
	entities.AddEntity(box2)

	return game.NewBasicSection(entities)
}

func CreateLevel1Section2(sharedEntities *entity.Entities) game.Section {
	entities := entity.NewEntities()
	entities.AddEntities(sharedEntities)

	ground := CreateGround()
	entities.AddEntity(ground)

	box1 := CreateBox(140, 140)
	entities.AddEntity(box1)

	box2 := CreateBox(160, 70)
	entities.AddEntity(box2)

	box3 := CreateBox(180, 50)
	entities.AddEntity(box3)

	return game.NewBasicSection(entities)
}
