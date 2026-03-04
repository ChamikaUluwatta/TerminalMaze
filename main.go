package main

import (
	"github.com/ChamikaUluwatta/TerminalMaze/generator"
	"github.com/ChamikaUluwatta/TerminalMaze/render"
)

func main() {
	maze := generator.MazeGenerator(20)
	render.RenderMaze(maze)
}
