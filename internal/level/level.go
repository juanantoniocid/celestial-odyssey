package level

import (
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/screen"
)

func LoadLevel1(player *entity.Player, character *entity.GameEntity, renderer screen.Renderer, inputHandler screen.InputHandler, collisionHandler screen.CollisionHandler, systemsInputHandler screen.SystemInputHandler) screen.Level {
	level1 := screen.NewLevel()

	level1.AddScenario(LoadLevel1Scenario1(player, character, renderer, inputHandler, collisionHandler, systemsInputHandler))
	level1.AddScenario(LoadLevel1Scenario2(player, character, renderer, inputHandler, collisionHandler, systemsInputHandler))

	return level1
}

func LoadLevel1Scenario1(player *entity.Player, character *entity.GameEntity, renderer screen.Renderer, inputHandler screen.InputHandler, collisionHandler screen.CollisionHandler, systemsInputHandler screen.SystemInputHandler) screen.Scenario {
	entities := entity.NewEntities()

	ground := entity.CreateGround()
	entities.AddEntity(ground)

	box1 := entity.CreateBox(100, 130)
	entities.AddEntity(box1)

	box2 := entity.CreateBox(120, 50)
	entities.AddEntity(box2)

	return screen.NewScenario(player, renderer, inputHandler, collisionHandler, systemsInputHandler, character, entities)
}

func LoadLevel1Scenario2(player *entity.Player, character *entity.GameEntity, renderer screen.Renderer, inputHandler screen.InputHandler, collisionHandler screen.CollisionHandler, systemsInputHandler screen.SystemInputHandler) screen.Scenario {
	entities := entity.NewEntities()

	ground := entity.CreateGround()
	entities.AddEntity(ground)

	box1 := entity.CreateBox(140, 140)
	entities.AddEntity(box1)

	box2 := entity.CreateBox(160, 70)
	entities.AddEntity(box2)

	box3 := entity.CreateBox(180, 50)
	entities.AddEntity(box3)

	return screen.NewScenario(player, renderer, inputHandler, collisionHandler, systemsInputHandler, character, entities)
}
