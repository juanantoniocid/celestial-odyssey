package factory

import (
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/screen"
	"celestial-odyssey/internal/system"
)

func LoadLevel1(character *entity.Entity, systems system.System, drawSystems system.DrawSystem) screen.Level {
	level1 := screen.NewLevel()

	level1.AddScenario(LoadLevel1Scenario1(character, systems, drawSystems))
	level1.AddScenario(LoadLevel1Scenario2(character, systems, drawSystems))

	return level1
}

func LoadLevel1Scenario1(character *entity.Entity, systems system.System, drawSystems system.DrawSystem) screen.Scenario {
	entities := entity.NewEntities()

	entities.AddEntity(character)

	ground := CreateGround()
	entities.AddEntity(ground)

	box1 := CreateBox(100, 130)
	entities.AddEntity(box1)

	box2 := CreateBox(120, 50)
	entities.AddEntity(box2)

	return screen.NewScenario(entities, systems, drawSystems)
}

func LoadLevel1Scenario2(character *entity.Entity, systems system.System, drawSystems system.DrawSystem) screen.Scenario {
	entities := entity.NewEntities()

	entities.AddEntity(character)

	ground := CreateGround()
	entities.AddEntity(ground)

	box1 := CreateBox(140, 140)
	entities.AddEntity(box1)

	box2 := CreateBox(160, 70)
	entities.AddEntity(box2)

	box3 := CreateBox(180, 50)
	entities.AddEntity(box3)

	return screen.NewScenario(entities, systems, drawSystems)
}
