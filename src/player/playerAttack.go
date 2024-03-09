package player

import (
	"github.com/Rayato159/nevers-space-shooting/src/world"
	"github.com/nsf/termbox-go"
)

var bulletPosition *PlayerPosition

func PlayerAttack(event termbox.Event) {
	bulletPosition = &PlayerPosition{X: playerPosition.X, Y: playerPosition.Y - 1}

	switch event.Type {
	case termbox.EventKey:
		switch event.Key {
		case termbox.KeySpace:
			bulletPosition.X = playerPosition.X
			bulletPosition.Y = playerPosition.Y - 1

			bulletFly('^')
			bulletConfine()
		}
	case termbox.EventError:
		panic(event.Err)
	}
}

func bulletFly(bullet rune) {
	world.World2D[bulletPosition.Y][bulletPosition.X] = bullet
	bulletPosition.Y--
}

func bulletConfine() {
	bulletPosition.Y = 1
	world.World2D[0][bulletPosition.X] = 'x'
}
