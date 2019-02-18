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
	seed = 126345
	percentage = 0.5

	board = Board.Create(x, y)
	board = Board.Randomize(board, seed, percentage)
	Board.Print(board)

	for {
		board = Board.Step(board)
		Board.Print(board)
		time.Sleep(delayms * time.Millisecond)
	}
}
