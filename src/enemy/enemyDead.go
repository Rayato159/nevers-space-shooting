package enemy

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/Rayato159/nevers-space-shooting/src/player"
	"github.com/Rayato159/nevers-space-shooting/src/world"
	"github.com/nsf/termbox-go"
)

func IsDead() {
	if player.BulletPosition == nil {
		return
	}

	if player.BulletPosition.X == enemyPosition.X {
		enemyPosition.X = rand.Intn(world.WIDTH - 2)
		world.World2D[enemyPosition.Y][enemyPosition.X] = 'E'

		if rand.NormFloat64() > player.PlayerCriticalChance {
			enemyHP.HP -= player.PlayerDamage * player.PlayerCriticalFactor
		} else {
			enemyHP.HP -= player.PlayerDamage
		}

		player.BulletPosition.X = 0
	}

	if enemyHP.HP == 0 {
		world.World2D[enemyPosition.Y][enemyPosition.X] = ' '

		fmt.Println("You Win!")

		termbox.Close()
		os.Exit(0)
	}
}
