package world

import "fmt"

const (
	WIDTH  = 10
	HEIGHT = 10
)

var world2D = [WIDTH + 2][HEIGHT + 2]rune{}

func Setup() {
	for y := range HEIGHT + 2 {
		setupHorizontalLine(y)

		for x := range WIDTH + 2 {
			setupVerticalLine(y, x)
		}
	}
}

func Draw() {
	for y := range HEIGHT + 2 {
		for x := range WIDTH + 2 {
			fmt.Print(string(world2D[y][x]))
		}

		fmt.Println()
	}
}

func setupHorizontalLine(y int) {
	if y == 0 || y == HEIGHT+2-1 {
		for x := range WIDTH + 2 {
			world2D[y][x] = 'x'
		}
	}
}

func setupVerticalLine(y, x int) {
	if y == 0 || y == HEIGHT+2-1 {
		return
	} else if x == 0 || x == WIDTH+2-1 {
		world2D[y][x] = 'x'
	} else {
		world2D[y][x] = ' '
	}
}
