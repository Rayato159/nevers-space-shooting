package player

import "github.com/Rayato159/nevers-space-shooting/src/world"

type (
	PlayerSprite struct {
		Sprite rune
	}

	PlayerHP struct {
		HP uint
	}
)

var (
	playerPosition               = &world.Vector2D{X: 1, Y: world.HEIGHT}
	playerSprite                 = &PlayerSprite{Sprite: 'P'}
	PlayerDamage         uint    = 25
	PlayerCriticalChance float64 = 0.3 // Percentage
	PlayerCriticalFactor uint    = 2   // Factor
)

func Setup() {
	world.World2D[playerPosition.Y][playerPosition.X] = playerSprite.Sprite
}
