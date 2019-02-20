package board

import (
	"fmt"
	"math/rand"
)

const (
	set   = "0"
	unset = " "
)

type Boardtype struct {
	boardX    int
	boardY    int
	iteration int
	board     [][]int
}

func (b *Boardtype) Create(x, y int) {
	b.boardX = x
	b.boardY = y
	b.board = make([][]int, y)
	for i := range b.board {
		b.board[i] = make([]int, x)
	}
}

func (b *Boardtype) Randomize(seed int64, percentage float64) {
	r := rand.New(rand.NewSource(seed))
	for y := range b.board {
		for x := range b.board[y] {
			if r.Intn(1000) > int(1000.0*percentage) {
				b.board[y][x] = 0
			} else {
				b.board[y][x] = 1
			}
		}
	}
}

func (b *Boardtype) Print() {
	for y := range b.board {
		for x := range b.board[y] {
			if b.board[y][x] == 0 {
				fmt.Printf("%s", unset)
			} else {
				fmt.Printf("%s", set)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Iteration: %d\n", b.iteration)
}

func (b *Boardtype) Step() {
	newBoard := make([][]int, len(b.board))
	for i := range b.board {
		newBoard[i] = make([]int, len(b.board[i]))
		copy(newBoard[i], b.board[i])
	}

	aliveNeighbors := 0
	for y := range b.board {
		for x := range b.board[y] {
			aliveNeighbors = b.getNeighbors(x, y)
			if b.board[y][x] == 1 && (aliveNeighbors == 2 || aliveNeighbors == 3) {
				newBoard[y][x] = 1
				// Rule 2.: Any live cell with two or three live neighbors lives on to the next generation.
			} else if b.board[y][x] == 0 && aliveNeighbors == 3 {
				// Rule 4.: Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
				newBoard[y][x] = 1
			} else if b.board[y][x] == 1 && (aliveNeighbors < 2 || aliveNeighbors > 3) {
				// Rule 1.: Any live cell with fewer than two live neighbors dies, as if by underpopulation.
				// Rule 3.: Any live cell with more than three live neighbors dies, as if by overpopulation.
				newBoard[y][x] = 0
			}
			aliveNeighbors = 0
		}
	}
	b.iteration++
	b.board = newBoard
}

func (b *Boardtype) getNeighbors(x, y int) int {
	aliveNeighbors := 0

	// Row above
	if x > 1 && y > 1 {
		if b.board[y-1][x-1] == 1 {
			aliveNeighbors++
		}
	}

	if y > 1 {
		if b.board[y-1][x] == 1 {
			aliveNeighbors++
		}
	}

	if x < b.boardX-1 && y > 1 {
		if b.board[y-1][x+1] == 1 {
			aliveNeighbors++
		}
	}

	// Same row
	if x > 1 {
		if b.board[y][x-1] == 1 {
			aliveNeighbors++
		}
	}

	if x < b.boardX-1 {
		if b.board[y][x+1] == 1 {
			aliveNeighbors++
		}
	}

	// Row below
	if x > 1 && y < b.boardY-1 {
		if b.board[y+1][x-1] == 1 {
			aliveNeighbors++
		}
	}

	if y < b.boardY-1 {
		if b.board[y+1][x] == 1 {
			aliveNeighbors++
		}
	}

	if x < b.boardX-1 && y < b.boardY-1 {
		if b.board[y+1][x+1] == 1 {
			aliveNeighbors++
		}
	}
	return aliveNeighbors
}
