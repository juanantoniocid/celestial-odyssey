package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/component"
	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/game"
	"celestial-odyssey/internal/graphics"
	"celestial-odyssey/internal/input"
	"celestial-odyssey/internal/screen"
	"celestial-odyssey/internal/systems"
)

func main() {
	cfg := config.LoadConfig()
	applyWindowSettings(cfg.Window)

	player := createPlayer(cfg.Player)
	renderer := createRenderer(cfg.Player, cfg.Screen, cfg.Ground)
	inputHandler := input.NewKeyboardHandler()
	physicsHandler := systems.NewPhysicsHandler()

	levels := createLevel(cfg.Screen, player, renderer, inputHandler, physicsHandler)
	screenManager := createScreenManager(cfg.Screen, []screen.Level{levels})

	g := createGame(screenManager)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func applyWindowSettings(cfg config.Window) {
	ebiten.SetWindowTitle(cfg.Title)
	ebiten.SetWindowSize(cfg.Dimensions.Width, cfg.Dimensions.Height)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
}

func createPlayer(cfg config.Player) (player *entity.Player) {
	player = entity.NewPlayer(cfg)

	return player
}

func createRenderer(cfgPlayer config.Player, cfgScreen config.Screen, cfgGround config.Ground) *graphics.Renderer {
	renderer := graphics.NewRenderer(cfgPlayer, cfgScreen, cfgGround)

	return renderer
}

func createLevel(cfg config.Screen, player *entity.Player, renderer screen.Renderer, inputHandler screen.InputHandler, physicsHandler screen.PhysicsHandler) screen.Level {
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

func createScreenManager(cfg config.Screen, levels []screen.Level) *screen.Manager {
	applyScreenSettings(cfg)
	manager := screen.NewManager(cfg)

	for _, level := range levels {
		manager.AddLevel(level)
	}

	return manager
}

func applyScreenSettings(cfg config.Screen) {
	ebiten.SetScreenClearedEveryFrame(cfg.ClearedEveryFrame)
}

func createGame(screenManager game.ScreenManager) *game.Game {
	g := game.NewGame(screenManager)
	g.Init()

	return g
}
