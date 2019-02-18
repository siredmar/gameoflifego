package board

import (
	"fmt"
	"math/rand"
)

const (
	set                   = "0"
	unset                 = " "
	lowerNeighborBoundary = 2
	upperNeighborBoundary = 3
)

var (
	boardX    int
	boardY    int
	iteration int
)

func Create(x, y int) [][]int {
	boardX = x
	boardY = y
	a := make([][]int, y)
	for i := range a {
		a[i] = make([]int, x)
	}
	return a
}

func Randomize(board [][]int, seed int64, percentage float32) [][]int {
	r := rand.New(rand.NewSource(seed))
	for y := range board {
		for x := range board[y] {
			if r.Intn(1000) > int(1000.0*percentage) {
				board[y][x] = 0
			} else {
				board[y][x] = 1
			}
		}
	}
	return board
}

func Print(board [][]int) {
	for y := range board {
		for x := range board[y] {
			if board[y][x] == 0 {
				fmt.Printf("%s", unset)
			} else {
				fmt.Printf("%s", set)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Iteration: %d\n", iteration)
	iteration++
}

func Step(board [][]int) [][]int {
	newBoard := make([][]int, len(board))
	for i := range board {
		newBoard[i] = make([]int, len(board[i]))
		copy(newBoard[i], board[i])
	}

	aliveNeighbors := 0
	for y := 0; y < boardY; y++ {
		for x := 0; x < boardX; x++ {
			aliveNeighbors = getNeighbors(board, x, y)
			if board[y][x] == 1 && (aliveNeighbors == 2 || aliveNeighbors == 3) {
				// Rule 2.: Any live cell with two or three live neighbors lives on to the next generation.
			} else if board[y][x] == 0 && aliveNeighbors == 3 {
				// Rule 4.: Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
				newBoard[y][x] = 1
			} else if board[y][x] == 1 && (aliveNeighbors < lowerNeighborBoundary || aliveNeighbors > upperNeighborBoundary) {
				// Rule 1.: Any live cell with fewer than two live neighbors dies, as if by underpopulation.
				// Rule 3.: Any live cell with more than three live neighbors dies, as if by overpopulation.
				newBoard[y][x] = 0
			}
			aliveNeighbors = 0
		}
	}
	return newBoard
}

func getNeighbors(b [][]int, x, y int) int {
	aliveNeighbors := 0

	// Row above
	if x > 1 && y > 1 {
		if b[y-1][x-1] == 1 {
			aliveNeighbors++
		}
	}

	if y > 1 {
		if b[y-1][x] == 1 {
			aliveNeighbors++
		}
	}

	if x < boardX-1 && y > 1 {
		if b[y-1][x+1] == 1 {
			aliveNeighbors++
		}
	}

	// Same row
	if x > 1 {
		if b[y][x-1] == 1 {
			aliveNeighbors++
		}
	}

	if x < boardX-1 {
		if b[y][x+1] == 1 {
			aliveNeighbors++
		}
	}

	// Row below
	if x > 1 && y < boardY-1 {
		if b[y+1][x-1] == 1 {
			aliveNeighbors++
		}
	}

	if y < boardY-1 {
		if b[y+1][x] == 1 {
			aliveNeighbors++
		}
	}

	if x < boardX-1 && y < boardY-1 {
		if b[y+1][x+1] == 1 {
			aliveNeighbors++
		}
	}
	return aliveNeighbors
}
