package factory

import (
	"celestial-odyssey/internal/entity"
	"celestial-odyssey/internal/game"
)

const (
	level1Section1ID = game.SectionID("level1-section1")
	level1Section2ID = game.SectionID("level1-section2")

	sideMargin = 8
)

func CreateLevel1(sharedEntities *entity.Entities) game.Level {
	level1 := game.NewBasicLevel()

	level1.AddSection(level1Section1ID, CreateLevel1Section1(sharedEntities))
	level1.AddSection(level1Section2ID, CreateLevel1Section2(sharedEntities))

	level1.AddTransition(game.NewTransition(level1Section1ID, level1Section2ID, playerIsOnRight, setPlayerToLeft))
	level1.AddTransition(game.NewTransition(level1Section2ID, level1Section1ID, playerIsOnLeft, setPlayerToRight))

	level1.SetCurrentSection(level1Section1ID)

	return level1
}

func CreateLevel1Section1(sharedEntities *entity.Entities) game.Section {
	entities := entity.NewEntities()
	entities.AddEntities(sharedEntities)

	ground := CreateGround()
	entities.AddEntity(ground)

	box1 := CreateBox(100, 130)
	entities.AddEntity(box1)

	box2 := CreateBox(120, 50)
	entities.AddEntity(box2)

	section := game.NewBasicSection(entities)
	return section
}

func CreateLevel1Section2(sharedEntities *entity.Entities) game.Section {
	entities := entity.NewEntities()
	entities.AddEntities(sharedEntities)

	ground := CreateGround()
	entities.AddEntity(ground)

	box1 := CreateBox(140, 140)
	entities.AddEntity(box1)

	box2 := CreateBox(160, 70)
	entities.AddEntity(box2)

	box3 := CreateBox(180, 50)
	entities.AddEntity(box3)

	section := game.NewBasicSection(entities)
	return section
}

func playerIsOnRight(entities *entity.Entities) bool {
	for _, e := range *entities {
		entityType, found := e.Type()
		if !found {
			continue
		}

		if entityType != entity.TypePlayer {
			continue
		}

		entityBounds, found := e.Bounds()
		if !found {
			continue
		}

		if entityBounds.Max.X >= 320 {
			return true
		}
	}

	return false
}

func setPlayerToRight(entities *entity.Entities) {
	for _, e := range *entities {
		entityType, found := e.Type()
		if !found {
			continue
		}

		if entityType != entity.TypePlayer {
			continue
		}

		entityPosition, found := e.Position()
		if !found {
			continue
		}

		entitySize, found := e.Size()
		if !found {
			continue
		}

		entityPosition.X = 320 - entitySize.Width - sideMargin
		e.SetPosition(entityPosition)
	}
}

func playerIsOnLeft(entities *entity.Entities) bool {
	for _, e := range *entities {
		entityType, found := e.Type()
		if !found {
			continue
		}

		if entityType != entity.TypePlayer {
			continue
		}

		entityBounds, found := e.Bounds()
		if !found {
			continue
		}

		if entityBounds.Min.X <= 0 {
			return true
		}
	}

	return false
}

func setPlayerToLeft(entities *entity.Entities) {
	for _, e := range *entities {
		entityType, found := e.Type()
		if !found {
			continue
		}

		if entityType != entity.TypePlayer {
			continue
		}

		entityPosition, found := e.Position()
		if !found {
			continue
		}

		entityPosition.X = sideMargin
		e.SetPosition(entityPosition)
	}
}
