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

	entities.AddGround()
	entities.AddBox(100, 130)
	entities.AddBox(120, 50)

	return screen.NewScenario(player, renderer, inputHandler, collisionHandler, systemsInputHandler, character, entities)
}

func LoadLevel1Scenario2(player *entity.Player, character *entity.GameEntity, renderer screen.Renderer, inputHandler screen.InputHandler, collisionHandler screen.CollisionHandler, systemsInputHandler screen.SystemInputHandler) screen.Scenario {
	entities := entity.NewEntities()

	entities.AddGround()
	entities.AddBox(140, 140)
	entities.AddBox(160, 70)
	entities.AddBox(180, 50)

	return screen.NewScenario(player, renderer, inputHandler, collisionHandler, systemsInputHandler, character, entities)
}
