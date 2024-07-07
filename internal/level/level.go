package level

import (
	"celestial-odyssey/internal/component"
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/screen"
)

func LoadLevel1(player *entity.Player, renderer screen.Renderer, inputHandler screen.InputHandler, physicsHandler screen.PhysicsHandler) screen.Level {
	level1 := screen.NewLevel()

	level1.AddScenario(LoadLevel1Scenario1(player, renderer, inputHandler, physicsHandler))
	level1.AddScenario(LoadLevel1Scenario2(player, renderer, inputHandler, physicsHandler))

	return level1
}

func LoadLevel1Scenario1(player *entity.Player, renderer screen.Renderer, inputHandler screen.InputHandler, physicsHandler screen.PhysicsHandler) screen.Scenario {
	entities := entity.NewEntities()

	box1 := entities.CreateEntity()
	box1.AddComponent("type", component.TypeBox)
	box1.AddComponent("position", &component.Position{X: 100, Y: 150})
	box1.AddComponent("size", &component.Size{Width: 100, Height: 22})

	box2 := entities.CreateEntity()
	box2.AddComponent("type", component.TypeBox)
	box2.AddComponent("position", &component.Position{X: 120, Y: 50})
	box2.AddComponent("size", &component.Size{Width: 80, Height: 50})

	ground := entities.CreateEntity()
	ground.AddComponent("type", component.TypeGround)
	ground.AddComponent("position", &component.Position{X: 0, Y: 172})
	ground.AddComponent("size", &component.Size{Width: 320, Height: 28})

	return screen.NewScenario(player, renderer, inputHandler, physicsHandler, entities)
}

func LoadLevel1Scenario2(player *entity.Player, renderer screen.Renderer, inputHandler screen.InputHandler, physicsHandler screen.PhysicsHandler) screen.Scenario {
	entities := entity.NewEntities()

	box1 := entities.CreateEntity()
	box1.AddComponent("type", component.TypeBox)
	box1.AddComponent("position", &component.Position{X: 100, Y: 150})
	box1.AddComponent("size", &component.Size{Width: 100, Height: 22})

	box2 := entities.CreateEntity()
	box2.AddComponent("type", component.TypeBox)
	box2.AddComponent("position", &component.Position{X: 120, Y: 50})
	box2.AddComponent("size", &component.Size{Width: 80, Height: 50})

	ground := entities.CreateEntity()
	ground.AddComponent("type", component.TypeGround)
	ground.AddComponent("position", &component.Position{X: 0, Y: 172})
	ground.AddComponent("size", &component.Size{Width: 320, Height: 28})

	return screen.NewScenario(player, renderer, inputHandler, physicsHandler, entities)
}
