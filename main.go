package main

import (
	"time"

	"github.com/Rayato159/nevers-space-shooting/src/enemy"
	"github.com/Rayato159/nevers-space-shooting/src/player"
	"github.com/Rayato159/nevers-space-shooting/src/systems"
	"github.com/Rayato159/nevers-space-shooting/src/world"
	"github.com/nsf/termbox-go"
)

const FPS = 10

var event termbox.Event

func main() {
	world.Setup()
	player.Setup()
	enemy.Setup()

	// Keyboard input detection
	termbox.Init()

	go eventLoopUpdate()

	for {
		systems.IsExit(event)

		world.Draw()
		world.ClearBullet()

		player.PlayerConfine()

		time.Sleep(1 / FPS * time.Second)
		systems.Render()
	}
}

func eventLoopUpdate() {
	for {
		event = termbox.PollEvent()

		player.Move(event)
		player.Attack(event)
	}
}
