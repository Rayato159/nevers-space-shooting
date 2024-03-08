package main

import (
	"github.com/Rayato159/nevers-space-shooting/src/player"
	"github.com/Rayato159/nevers-space-shooting/src/systems"
	"github.com/Rayato159/nevers-space-shooting/src/world"
)

func main() {
	world.Setup()
	player.Setup()

	for {
		world.Draw()

		systems.Render()
	}
}
