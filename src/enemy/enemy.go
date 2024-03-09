package enemy

import (
	"fmt"

	"github.com/Rayato159/nevers-space-shooting/src/world"
)

type (
	EnemySprite struct {
		Sprite rune
	}

	EnemyHP struct {
		HP uint
	}
)

var (
	enemyPosition = &world.Vector2D{X: world.WIDTH, Y: 1}
	enemySprite   = &EnemySprite{Sprite: 'E'}
	enemyHP       = &EnemyHP{HP: 100}
)

func Setup() {
	world.World2D[enemyPosition.Y][enemyPosition.X] = enemySprite.Sprite
}

func ShowEnemyHP() {
	fmt.Printf("Enemy HP: %d\n", enemyHP.HP)
}
