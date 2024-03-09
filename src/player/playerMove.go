package player

import (
	"github.com/Rayato159/nevers-space-shooting/src/world"
	"github.com/nsf/termbox-go"
)

func Move(event termbox.Event) {
	switch event.Type {
	case termbox.EventKey:
		switch event.Key {
		case termbox.KeyArrowUp:
			termbox.Sync()
			world.World2D[playerPosition.Y][playerPosition.X] = ' '
			playerPosition.Y--
			world.World2D[playerPosition.Y][playerPosition.X] = playerSprite.Sprite
		case termbox.KeyArrowRight:
			termbox.Sync()
			world.World2D[playerPosition.Y][playerPosition.X] = ' '
			playerPosition.X++
			world.World2D[playerPosition.Y][playerPosition.X] = playerSprite.Sprite
		case termbox.KeyArrowDown:
			termbox.Sync()
			world.World2D[playerPosition.Y][playerPosition.X] = ' '
			playerPosition.Y++
			world.World2D[playerPosition.Y][playerPosition.X] = playerSprite.Sprite
		case termbox.KeyArrowLeft:
			termbox.Sync()
			world.World2D[playerPosition.Y][playerPosition.X] = ' '
			playerPosition.X--
			world.World2D[playerPosition.Y][playerPosition.X] = playerSprite.Sprite
		}
	case termbox.EventError:
		panic(event.Err)
	}
}
