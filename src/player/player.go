package player

import "github.com/Rayato159/nevers-space-shooting/src/world"

type (
	PlayerPosition struct {
		X, Y int
	}

	PlayerSprite struct {
		Sprite rune
	}

	PlayerHP struct {
		HP uint
	}
)

var (
	playerPosition = &PlayerPosition{X: 1, Y: world.HEIGHT}
	playerSprite   = &PlayerSprite{Sprite: 'P'}
	playerHP       = &PlayerHP{HP: 100}
)

func Setup() {
	world.World2D[playerPosition.Y][playerPosition.X] = playerSprite.Sprite
}
