package level

import (
	"celestial-odyssey/internal/component"
	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/screen"
)

func LoadLevel(cfg config.Screen, player *entity.Player, renderer screen.Renderer, inputHandler screen.InputHandler, physicsHandler screen.PhysicsHandler) screen.Level {
	level := screen.NewLevel()

	entityManager := entity.NewEntityManager()

	box1 := entityManager.CreateEntity()
	box1.AddComponent("type", component.TypeBox)
	box1.AddComponent("position", &component.Position{X: 100, Y: 150})
	box1.AddComponent("size", &component.Size{Width: 100, Height: 22})

	box2 := entityManager.CreateEntity()
	box2.AddComponent("type", component.TypeBox)
	box2.AddComponent("position", &component.Position{X: 120, Y: 50})
	box2.AddComponent("size", &component.Size{Width: 80, Height: 50})

	ground := entityManager.CreateEntity()
	ground.AddComponent("type", component.TypeGround)
	ground.AddComponent("position", &component.Position{X: 0, Y: 172})
	ground.AddComponent("size", &component.Size{Width: 320, Height: 28})

	landingSite := screen.NewScenario(player, renderer, inputHandler, physicsHandler, entityManager, cfg.Dimensions.Width, cfg.Dimensions.Height)
	sandDunes := screen.NewScenario(player, renderer, inputHandler, physicsHandler, entity.NewEntityManager(), cfg.Dimensions.Width, cfg.Dimensions.Height)
	ruinedTemple := screen.NewScenario(player, renderer, inputHandler, physicsHandler, entity.NewEntityManager(), cfg.Dimensions.Width, cfg.Dimensions.Height)

	level.AddScenario(landingSite)
	level.AddScenario(sandDunes)
	level.AddScenario(ruinedTemple)

	return level
}
