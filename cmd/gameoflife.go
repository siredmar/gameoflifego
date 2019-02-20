package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	Board "github.com/siredmar/gameoflifego/internal/board"
)

var (
	x          *int
	y          *int
	delayms    *int
	seed       *int64
	percentage *float64
	board      [][]int
)

var Usage = func() {
	desc := `This is 'Conway's Game of Life' written in golang.
It follows the rules:
1. Any live cell with fewer than two live neighbors dies, 
as if by underpopulation.
2. Any live cell with two or three live neighbors lives on 
to the next generation.
3. Any live cell with more than three live neighbors dies, 
as if by overpopulation.
4. Any dead cell with exactly three live neighbors becomes 
a live cell, as if by reproduction.
	
The Pattern will be generated randomly`

	fmt.Fprintf(os.Stderr, "%s\n\n", desc)
	fmt.Fprintf(os.Stderr, "Usage:\n")

	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	x = flag.Int("columns", 80, "Columns")
	y = flag.Int("rows", 22, "Rows")
	percentage = flag.Float64("density", 0.5, "Random density of living cells")
	delayms = flag.Int("delay", 200, "Delay in milliseconds between iterations")
	seed = flag.Int64("seed", 1234, "Seed for generating a random board")
	flag.Parse()

	board := Board.Boardtype{}

	board.Create(*x, *y)
	board.Randomize(*seed, *percentage)
	board.Print()

	for {
		board.Step()
		board.Print()
		time.Sleep(time.Duration(*delayms) * time.Millisecond)
	}
}
