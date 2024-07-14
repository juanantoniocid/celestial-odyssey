package factory

import (
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/graphics"
	"celestial-odyssey/internal/screen"
	"celestial-odyssey/internal/system"
)

func LoadLevel1(player *entity.Player, character *entity.Entity, renderer screen.Renderer, inputHandler screen.InputHandler, collisionHandler screen.CollisionHandler, systems system.System, drawSystems graphics.DrawSystem) screen.Level {
	level1 := screen.NewLevel()

	level1.AddScenario(LoadLevel1Scenario1(player, character, renderer, inputHandler, collisionHandler, systems, drawSystems))
	level1.AddScenario(LoadLevel1Scenario2(player, character, renderer, inputHandler, collisionHandler, systems, drawSystems))

	return level1
}

func LoadLevel1Scenario1(player *entity.Player, character *entity.Entity, renderer screen.Renderer, inputHandler screen.InputHandler, collisionHandler screen.CollisionHandler, systems system.System, drawSystems graphics.DrawSystem) screen.Scenario {
	entities := entity.NewEntities()

	entities.AddEntity(character)

	ground := CreateGround()
	entities.AddEntity(ground)

	box1 := CreateBox(100, 130)
	entities.AddEntity(box1)

	box2 := CreateBox(120, 50)
	entities.AddEntity(box2)

	return screen.NewScenario(player, renderer, inputHandler, collisionHandler, entities, systems, drawSystems)
}

func LoadLevel1Scenario2(player *entity.Player, character *entity.Entity, renderer screen.Renderer, inputHandler screen.InputHandler, collisionHandler screen.CollisionHandler, systems system.System, drawSystems graphics.DrawSystem) screen.Scenario {
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

	return screen.NewScenario(player, renderer, inputHandler, collisionHandler, entities, systems, drawSystems)
}
