package player

import (
	"github.com/Rayato159/nevers-space-shooting/src/world"
	"github.com/nsf/termbox-go"
)

var bulletPosition *world.Vector2D

func Attack(event termbox.Event) {
	bulletPosition = &world.Vector2D{X: playerPosition.X, Y: playerPosition.Y - 1}

	switch event.Type {
	case termbox.EventKey:
		switch event.Key {
		case termbox.KeySpace:
			bulletPosition.X = playerPosition.X
			bulletPosition.Y = playerPosition.Y - 1

			bulletFly()
		}
	case termbox.EventError:
		panic(event.Err)
	}
}

func bulletFly() {
	for bulletPosition.Y > 0 {
		world.World2D[bulletPosition.Y][bulletPosition.X] = '|'
		bulletPosition.Y--
	}
}
