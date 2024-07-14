package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/factory"
	"celestial-odyssey/internal/game"
	"celestial-odyssey/internal/graphics"
	"celestial-odyssey/internal/input"
	"celestial-odyssey/internal/screen"
	"celestial-odyssey/internal/system"
)

func main() {
	cfg := config.LoadConfig()
	applyWindowSettings(cfg.Window)

	player := createPlayer(cfg.Player)
	character := factory.CreatePlayer()
	renderer := createRenderer(cfg.Player, cfg.Screen, cfg.Ground)
	inputHandler := input.NewKeyboardHandler()
	collisionHandler := system.NewCollisionHandler()

	inputManager := system.NewInput()
	movementManager := system.NewMovement()
	systemManager := system.NewManager(inputManager, inputManager, movementManager, collisionHandler)

	levels := factory.LoadLevel1(player, character, renderer, inputHandler, collisionHandler, systemManager)
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

func createScreenManager(cfg config.Screen, levels []screen.Level) *screen.Manager {
	applyScreenSettings(cfg)
	manager := screen.NewManager(cfg)

	for _, l := range levels {
		manager.AddLevel(l)
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
