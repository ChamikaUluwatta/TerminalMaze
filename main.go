package main

import (
	"github.com/ChamikaUluwatta/TerminalMaze/generator"
	"github.com/ChamikaUluwatta/TerminalMaze/render"
)

func main() {
	maze := generator.MazeGenerator(4)
	render.RenderMaze(maze)
}
