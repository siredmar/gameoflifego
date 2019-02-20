package main

import (
	"time"

	Board "github.com/siredmar/gameoflifego/internal/board"
)

var (
	x          int
	y          int
	delayms    time.Duration
	board      [][]int
	seed       int64
	percentage float32
)

func main() {
	x = 80
	y = 22
	delayms = 200
	seed = 1234
	percentage = 0.5

	board := Board.Boardtype{}

	board.Create(x, y)
	board.Randomize(seed, percentage)
	board.Print()

	for {
		board.Step()
		board.Print()
		time.Sleep(delayms * time.Millisecond)
	}
}
