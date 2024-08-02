package factory

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"celestial-odyssey/internal/component"
	"celestial-odyssey/internal/config"
	"celestial-odyssey/internal/entity"
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
func CreatePlayer() *entity.Entity {
	player := entity.NewEntity()

	player.SetType(entity.TypePlayer)
	player.SetPosition(component.Position{X: 8, Y: 140})
	player.SetSize(component.Size{Width: 16, Height: 32})
	player.SetVelocity(component.Velocity{VX: 0, VY: 0})
	player.SetAction(component.Action{Left: false, Right: false, Jump: false})
	player.SetInput(component.Input{Left: ebiten.KeyLeft, Right: ebiten.KeyRight, Jump: ebiten.KeySpace})
	player.SetColor(component.Color{Color: color.RGBA{R: 255, G: 255, B: 255, A: 255}})

	return player
}

// CreateBox creates a box entity.
func CreateBox(x, y float64) *entity.Entity {
	box := entity.NewEntity()

	box.SetType(entity.TypeBox)
	box.SetPosition(component.Position{X: x, Y: y})
	box.SetSize(component.Size{Width: boxWidth, Height: boxHeight})
	box.SetColor(component.Color{Color: color.RGBA{R: 255, G: 165, B: 0, A: 255}})

	return box
}

// CreateGround creates a ground entity.
func CreateGround() *entity.Entity {
	ground := entity.NewEntity()

	ground.SetType(entity.TypeGround)
	ground.SetPosition(component.Position{X: groundPositionX, Y: groundPositionY})
	ground.SetSize(component.Size{Width: groundWidth, Height: groundHeight})
	ground.SetColor(component.Color{Color: color.RGBA{R: 169, G: 169, B: 169, A: 255}})

	return ground
}
