package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

const (
	w, h = 60, 20
	// default is 1ms, "speed" is a factor to increase it.
	speed = 100
)

func inverDir(dir int) int {
	if dir == 0 {
		return 1
	}
	return 0

}

func main() {

	// initialize the board
	board := make([][]bool, w)

	// initializes inner slices (board height)
	for i := range board {
		board[i] = make([]bool, h)
	}

	countw := 0
	counth := 0
	// 0 means down|right, 1 means up|left
	dirw := 0
	dirh := 0

	for {
		screen.Clear()
		screen.MoveTopLeft()

		//Draw the board
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				if !board[x][y] {
					fmt.Print("  ")
				} else {
					fmt.Print("🏐")
				}
				//fmt.Print(board[x][y])
			}
			fmt.Println()
		}

		// set previous position to false

		board[countw][counth] = false

		// go to the next cell
		if dirw == 0 {
			countw++
		} else {
			countw--
		}
		if dirh == 0 {
			counth++
		} else {
			counth--
		}

		// check if we are in the edges
		if countw == w-1 || countw == 0 {
			dirw = inverDir(dirw)
		}
		if counth == h-1 || counth == 0 {
			dirh = inverDir(dirh)
		}

		fmt.Println(countw, counth)
		time.Sleep(time.Millisecond * speed)
		board[countw][counth] = true

	}
}
