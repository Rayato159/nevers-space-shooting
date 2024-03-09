package player

import (
	"github.com/Rayato159/nevers-space-shooting/src/world"
	"github.com/nsf/termbox-go"
)

var BulletPosition *world.Vector2D

func Attack(event termbox.Event) {
	BulletPosition = &world.Vector2D{X: playerPosition.X, Y: playerPosition.Y - 1}

	switch event.Type {
	case termbox.EventKey:
		switch event.Key {
		case termbox.KeySpace:
			BulletPosition.X = playerPosition.X
			BulletPosition.Y = playerPosition.Y - 1

			bulletFly()
		}
	case termbox.EventError:
		panic(event.Err)
	}
}

func bulletFly() {
	for BulletPosition.Y > 0 {
		world.World2D[BulletPosition.Y][BulletPosition.X] = '|'
		BulletPosition.Y--
	}
}
