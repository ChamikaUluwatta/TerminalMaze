package render

import (
	"fmt"

	"github.com/ChamikaUluwatta/TerminalMaze/generator"
)

func RenderMaze(maze generator.Maze) {
	size := maze.Size
	rows := 2*size + 1
	cols := 2*size + 1

	grid := make([][]rune, rows)
	for i := range grid {
		grid[i] = make([]rune, cols)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	for i := 0; i < rows; i += 2 {
		for j := 0; j < cols; j += 2 {
			grid[i][j] = '+'
		}
	}

	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			cell := maze.Grid[r][c]
			if cell.Up {
				grid[2*r][2*c+1] = '-'
			}
			if cell.Down {
				grid[2*r+2][2*c+1] = '-'
			}
			if cell.Left {
				grid[2*r+1][2*c] = '|'
			}
			if cell.Right {
				grid[2*r+1][2*c+2] = '|'
			}
		}
	}

	grid[0][2*maze.Start[1]+1] = ' '
	grid[2*maze.End[0]+2][2*maze.End[1]+1] = ' '

	for _, row := range grid {
		fmt.Println(string(row))
	}
}
