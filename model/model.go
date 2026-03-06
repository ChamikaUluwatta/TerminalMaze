package model

import (
	tea "charm.land/bubbletea/v2"
	"github.com/ChamikaUluwatta/TerminalMaze/generator"
	"github.com/ChamikaUluwatta/TerminalMaze/render"
)

type Model struct {
	maze generator.Maze
	size int
}

func InitialModel(size int) Model {
	return Model{
		maze: generator.MazeGenerator(size),
		size: size,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			m.maze = generator.MazeGenerator(m.size)
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() tea.View {
	s := render.RenderMaze(m.maze)
	s += "\n[r] regenerate  [q] quit\n"
	return tea.NewView(s)
}
