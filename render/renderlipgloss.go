package render

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/ChamikaUluwatta/TerminalMaze/generator"
)

var (
	wallStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#bd93f9")).
			Bold(true)
	playerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ff5555")).
			Bold(true)
)

func RenderMazeWithLipgloss(maze generator.Maze, playerRow, playerCol, zoom int) string {
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

	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			cell := maze.Grid[r][c]
			if cell.Up {
				grid[2*r][2*c+1] = '\u2500' // ─
			}
			if cell.Down {
				grid[2*r+2][2*c+1] = '\u2500' // ─
			}
			if cell.Left {
				grid[2*r+1][2*c] = '\u2502' // │
			}
			if cell.Right {
				grid[2*r+1][2*c+2] = '\u2502' // │
			}
		}
	}

	for i := 0; i < rows; i += 2 {
		for j := 0; j < cols; j += 2 {
			up := i > 0 && grid[i-1][j] == '\u2502'
			down := i < rows-1 && grid[i+1][j] == '\u2502'
			left := j > 0 && grid[i][j-1] == '\u2500'
			right := j < cols-1 && grid[i][j+1] == '\u2500'
			grid[i][j] = cornerChar(up, down, left, right)
		}
	}

	grid[0][2*maze.Start[1]+1] = ' '
	grid[2*maze.End[0]+2][2*maze.End[1]+1] = ' '

	playerGridRow := 2*playerRow + 1
	playerGridCol := 2*playerCol + 1

	repeatH := func(ch rune, odd bool) string {
		if odd {
			return strings.Repeat(string(ch), zoom)
		}
		return string(ch)
	}

	playerDot := "●" + strings.Repeat(" ", zoom-1)

	var sb strings.Builder
	for i, row := range grid {
		vRepeat := 1
		if i%2 == 1 && zoom >= 3 {
			vRepeat = zoom / 2
		}
		for v := 0; v < vRepeat; v++ {
			if i == playerGridRow && v == vRepeat/2 {
				var before, after strings.Builder
				for j, ch := range row {
					if j < playerGridCol {
						before.WriteString(repeatH(ch, j%2 == 1))
					} else if j > playerGridCol {
						after.WriteString(repeatH(ch, j%2 == 1))
					}
				}
				sb.WriteString(wallStyle.Render(before.String()))
				sb.WriteString(playerStyle.Render(playerDot))
				sb.WriteString(wallStyle.Render(after.String()))
			} else {
				var line strings.Builder
				for j, ch := range row {
					line.WriteString(repeatH(ch, j%2 == 1))
				}
				sb.WriteString(wallStyle.Render(line.String()))
			}
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}

func cornerChar(up, down, left, right bool) rune {
	switch {
	case up && down && left && right:
		return '\u253C' // ┼
	case up && down && left:
		return '\u2524' // ┤
	case up && down && right:
		return '\u251C' // ├
	case up && left && right:
		return '\u2534' // ┴
	case down && left && right:
		return '\u252C' // ┬
	case up && down:
		return '\u2502' // │
	case left && right:
		return '\u2500' // ─
	case down && right:
		return '\u250C' // ┌
	case down && left:
		return '\u2510' // ┐
	case up && right:
		return '\u2514' // └
	case up && left:
		return '\u2518' // ┘
	case up:
		return '\u2575' // ╵
	case down:
		return '\u2577' // ╷
	case left:
		return '\u2574' // ╴
	case right:
		return '\u2576' // ╶
	default:
		return ' '
	}
}
