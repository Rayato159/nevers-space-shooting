package world

import "fmt"

const (
	WIDTH  = 10
	HEIGHT = 10
)

func Draw() {
	for y := range HEIGHT + 2 {
		drawHorizontalLine(y)

		for x := range WIDTH + 2 {
			drawVerticalLine(x, y)
		}

		fmt.Println()
	}
}

func drawHorizontalLine(y int) {
	if y == 0 || y == HEIGHT+2-1 {
		for range WIDTH + 2 {
			fmt.Print("x")
		}
	}
}

func drawVerticalLine(x, y int) {
	if (x == 0 || x == WIDTH+2-1) && (y != 0 && y != HEIGHT+2-1) {
		fmt.Print("x")
	} else {
		fmt.Print(" ")
	}
}
