package main

import (
	"time"

	"github.com/Rayato159/nevers-space-shooting/src/player"
	"github.com/Rayato159/nevers-space-shooting/src/systems"
	"github.com/Rayato159/nevers-space-shooting/src/world"
	"github.com/nsf/termbox-go"
)

const FPS = 10

func main() {
	world.Setup()
	player.Setup()

	// Keyboard input detection
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	for {
		world.Draw()

		event := termbox.PollEvent()
		player.Move(event)

		time.Sleep(1 / FPS * time.Second)

		systems.Render()
	}
}
