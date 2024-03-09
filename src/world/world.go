package world

import (
	"fmt"
)

type Vector2D struct {
	X, Y int
}

const (
	WIDTH  = 10
	HEIGHT = 10
)

var World2D = [WIDTH + 2][HEIGHT + 2]rune{}

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
			fmt.Print(string(World2D[y][x]))
		}

		fmt.Println()
	}
}

func ClearBullet() {
	for y := 0; y < HEIGHT+1; y++ {
		for x := 0; x < WIDTH+1; x++ {
			if World2D[y][x] == '^' {
				World2D[y][x] = ' '
			}
		}
	}
}

func setupHorizontalLine(y int) {
	if y == 0 || y == HEIGHT+2-1 {
		for x := range WIDTH + 2 {
			World2D[y][x] = 'x'
		}
	}
}

func setupVerticalLine(y, x int) {
	if y == 0 || y == HEIGHT+2-1 {
		return
	} else if x == 0 || x == WIDTH+2-1 {
		World2D[y][x] = 'x'
	} else {
		World2D[y][x] = ' '
	}
}
