package main

import (
	"github.com/Rayato159/nevers-space-shooting/src/systems"
	"github.com/Rayato159/nevers-space-shooting/src/world"
)

func main() {
	world.Setup()

	for {
		world.Draw()

		systems.Render()
	}
}
