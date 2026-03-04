package generator

import (
	"math/rand"
)

type Cell struct {
	visited bool
	left    bool
	right   bool
	up      bool
	down    bool
}

func cellGenerator(size int) [][]Cell {

	twoDArray := make([][]Cell, size)

	for i := range twoDArray {
		twoDArray[i] = make([]Cell, size)
		for j := range twoDArray[i] {
			twoDArray[i][j] = Cell{
				left:  true,
				right: true,
				up:    true,
				down:  true,
			}
		}
	}
	return twoDArray
}

func mazeGenerator(size int) [][]Cell {
	grid := cellGenerator(size)
	start := rand.Intn(size)
	cellStart := grid[0][start]
	dfs(&grid)
	return grid
}

func dfs(grid *[][]Cell) {

}
