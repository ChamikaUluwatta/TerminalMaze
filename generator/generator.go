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

type Point struct {
	x int
	y int
}

type Maze struct {
	Grid  [][]Cell
	Size  int
	Start [2]int
	End   [2]int
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

func MazeGenerator(size int) Maze {
	grid := cellGenerator(size)
	start := rand.Intn(size)
	traversalDfs(&grid, 0, start)
	end := findingEndBFS(&grid, start)

	return Maze{
		Grid:  grid,
		Size:  size,
		Start: [2]int{0, start},
		End:   [2]int{end.x, end.y},
	}
}

func findingEndBFS(grid *[][]Cell, start int) Point {
	size := len(*grid)
	queue := []Point{{0, start}}
	visited := make(map[Point]bool)
	distance := make(map[Point]int)
	visited[Point{0, start}] = true
	distance[Point{0, start}] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		passableNeighbors := getPassableNeighbors(grid, current.x, current.y)

		for _, neighbor := range passableNeighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				distance[neighbor] = distance[current] + 1
				queue = append(queue, neighbor)
			}
		}
	}

	// Find the farthest reachable cell on the last row
	best := Point{size - 1, 0}
	bestDist := -1
	for col := 0; col < size; col++ {
		p := Point{size - 1, col}
		if d, ok := distance[p]; ok && d > bestDist {
			bestDist = d
			best = p
		}
	}
	return best
}

func traversalDfs(grid *[][]Cell, x, y int) {
	stack := []Point{{x, y}}
	(*grid)[x][y].visited = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		neighbors := getUnvisitedNeighbors(grid, current.x, current.y)

		if len(neighbors) == 0 {
			stack = stack[:len(stack)-1]
		} else {
			randIndex := rand.Intn(len(neighbors))
			chosen := neighbors[randIndex]
			removeWall(grid, current, chosen)
			(*grid)[chosen.x][chosen.y].visited = true
			stack = append(stack, chosen)
		}
	}
}

func getPassableNeighbors(grid *[][]Cell, x, y int) []Point {
	var neighbors []Point
	if x > 0 && !(*grid)[x][y].up {
		neighbors = append(neighbors, Point{x - 1, y})
	}
	if x < len(*grid)-1 && !(*grid)[x][y].down {
		neighbors = append(neighbors, Point{x + 1, y})
	}
	if y > 0 && !(*grid)[x][y].left {
		neighbors = append(neighbors, Point{x, y - 1})
	}
	if y < len((*grid)[0])-1 && !(*grid)[x][y].right {
		neighbors = append(neighbors, Point{x, y + 1})
	}
	return neighbors
}

func getUnvisitedNeighbors(grid *[][]Cell, x, y int) []Point {
	var neighbors []Point
	if x > 0 && !(*grid)[x-1][y].visited {
		neighbors = append(neighbors, Point{x - 1, y})
	}
	if x < len(*grid)-1 && !(*grid)[x+1][y].visited {
		neighbors = append(neighbors, Point{x + 1, y})
	}
	if y > 0 && !(*grid)[x][y-1].visited {
		neighbors = append(neighbors, Point{x, y - 1})
	}
	if y < len((*grid)[0])-1 && !(*grid)[x][y+1].visited {
		neighbors = append(neighbors, Point{x, y + 1})
	}
	return neighbors
}

func removeWall(grid *[][]Cell, current Point, next Point) {
	//x+1 = down, x-1 = up, y+1 = right, y-1 = left

	if current.x+1 == next.x {
		(*grid)[current.x][current.y].down = false
		(*grid)[next.x][next.y].up = false
	} else if current.x-1 == next.x {
		(*grid)[current.x][current.y].up = false
		(*grid)[next.x][next.y].down = false
	} else if current.y+1 == next.y {
		(*grid)[current.x][current.y].right = false
		(*grid)[next.x][next.y].left = false
	} else if current.y-1 == next.y {
		(*grid)[current.x][current.y].left = false
		(*grid)[next.x][next.y].right = false
	}
}
