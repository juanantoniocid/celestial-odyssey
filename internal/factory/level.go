package factory

import (
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/game"
	"celestial-odyssey/internal/system/behavior"
	"celestial-odyssey/internal/system/graphics"
)

func CreateLevel1(sharedEntities *entity.Entities, updateSystem behavior.UpdateSystem, renderer graphics.Renderer) game.Level {
	level1 := game.NewLevel()

	level1.AddScenario(CreateLevel1Scenario1(sharedEntities, updateSystem, renderer))
	level1.AddScenario(CreateLevel1Scenario2(sharedEntities, updateSystem, renderer))

	return level1
}

func CreateLevel1Scenario1(sharedEntities *entity.Entities, updateSystem behavior.UpdateSystem, renderer graphics.Renderer) game.Scenario {
	entities := entity.NewEntities()
	entities.AddEntities(sharedEntities)

	ground := CreateGround()
	entities.AddEntity(ground)

	box1 := CreateBox(100, 130)
	entities.AddEntity(box1)

	box2 := CreateBox(120, 50)
	entities.AddEntity(box2)

	return game.NewScenario(entities, updateSystem, renderer)
}

func CreateLevel1Scenario2(sharedEntities *entity.Entities, updateSystem behavior.UpdateSystem, renderer graphics.Renderer) game.Scenario {
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

	return game.NewScenario(entities, updateSystem, renderer)
}
