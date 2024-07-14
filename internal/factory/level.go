package factory

import (
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/screen"
)

func LoadLevel1(player *entity.Player, character *entity.GameEntity, renderer screen.Renderer, inputHandler screen.InputHandler, collisionHandler screen.CollisionHandler, systemManager screen.SystemManager) screen.Level {
	level1 := screen.NewLevel()

	level1.AddScenario(LoadLevel1Scenario1(player, character, renderer, inputHandler, collisionHandler, systemManager))
	level1.AddScenario(LoadLevel1Scenario2(player, character, renderer, inputHandler, collisionHandler, systemManager))

	return level1
}

func LoadLevel1Scenario1(player *entity.Player, character *entity.GameEntity, renderer screen.Renderer, inputHandler screen.InputHandler, collisionHandler screen.CollisionHandler, systemManager screen.SystemManager) screen.Scenario {
	collection := entity.NewCollection()

	collection.AddGameEntity(character)

	ground := CreateGround()
	collection.AddGameEntity(ground)

	box1 := CreateBox(100, 130)
	collection.AddGameEntity(box1)

	box2 := CreateBox(120, 50)
	collection.AddGameEntity(box2)

	return screen.NewScenario(player, renderer, inputHandler, collisionHandler, collection, systemManager)
}

func LoadLevel1Scenario2(player *entity.Player, character *entity.GameEntity, renderer screen.Renderer, inputHandler screen.InputHandler, collisionHandler screen.CollisionHandler, systemManager screen.SystemManager) screen.Scenario {
	collection := entity.NewCollection()

	collection.AddGameEntity(character)

	ground := CreateGround()
	collection.AddGameEntity(ground)

	box1 := CreateBox(140, 140)
	collection.AddGameEntity(box1)

	box2 := CreateBox(160, 70)
	collection.AddGameEntity(box2)

	box3 := CreateBox(180, 50)
	collection.AddGameEntity(box3)

	return screen.NewScenario(player, renderer, inputHandler, collisionHandler, collection, systemManager)
}
