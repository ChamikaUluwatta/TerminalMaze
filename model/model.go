package model

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/ChamikaUluwatta/TerminalMaze/generator"
	"github.com/ChamikaUluwatta/TerminalMaze/render"
)

type Model struct {
	maze   generator.Maze
	size   int
	width  int
	height int
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
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
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
	maze := render.RenderMazeWithLipgloss(m.maze)
	help := "[r] regenerate  [q] quit"
	content := lipgloss.JoinVertical(lipgloss.Center, maze, "", help)
	s := lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
	return tea.NewView(s)
}
