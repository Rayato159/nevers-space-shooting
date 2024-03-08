package player

import "github.com/Rayato159/nevers-space-shooting/src/world"

func PlayerConfine() {
	if playerPosition.X < 1 {
		playerPosition.X = 1

		world.World2D[playerPosition.Y][0] = 'x'
		world.World2D[playerPosition.Y][1] = playerSprite.Sprite

	} else if playerPosition.X > world.WIDTH {
		playerPosition.X = world.WIDTH

		world.World2D[playerPosition.Y][world.WIDTH+1] = 'x'
		world.World2D[playerPosition.Y][world.WIDTH] = playerSprite.Sprite
	}

	if playerPosition.Y < 1 {
		playerPosition.Y = 1

		world.World2D[0][playerPosition.X] = 'x'
		world.World2D[1][playerPosition.X] = playerSprite.Sprite

	} else if playerPosition.Y > world.HEIGHT {
		playerPosition.Y = world.HEIGHT

		world.World2D[world.HEIGHT+1][playerPosition.X] = 'x'
		world.World2D[world.HEIGHT][playerPosition.X] = playerSprite.Sprite
	}
}
