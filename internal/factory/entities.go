package factory

import (
	"celestial-odyssey/internal/entity"
	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/component"
	"celestial-odyssey/internal/config"
)

const (
	boxWidth  = 30
	boxHeight = 30

	groundPositionX = 0
	groundPositionY = 172

	groundWidth  = config.ScreenWidth
	groundHeight = 28
)

// CreatePlayer creates a player entity.
func CreatePlayer() *entity.GameEntity {
	player := entity.NewGameEntity()

	player.SetType(entity.TypePlayer)
	player.SetPosition(component.Position{X: 0, Y: 0})
	player.SetSize(component.Size{Width: 20, Height: 40})
	player.SetVelocity(component.Velocity{VX: 0, VY: 0})
	player.SetInput(component.Input{Left: false, Right: false, Jump: false})
	player.SetInputMap(component.InputMap{Left: ebiten.KeyLeft, Right: ebiten.KeyRight, Jump: ebiten.KeySpace})

	return player
}

// CreateBox creates a box entity.
func CreateBox(x, y float64) *entity.GameEntity {
	box := entity.NewGameEntity()

	box.SetType(entity.TypeBox)
	box.SetPosition(component.Position{X: x, Y: y})
	box.SetSize(component.Size{Width: boxWidth, Height: boxHeight})

	return box
}

// CreateGround creates a ground entity.
func CreateGround() *entity.GameEntity {
	ground := entity.NewGameEntity()

	ground.SetType(entity.TypeGround)
	ground.SetPosition(component.Position{X: groundPositionX, Y: groundPositionY})
	ground.SetSize(component.Size{Width: groundWidth, Height: groundHeight})

	return ground
}
