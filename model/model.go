package model

import (
	"strconv"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/ChamikaUluwatta/TerminalMaze/generator"
	"github.com/ChamikaUluwatta/TerminalMaze/render"
)

type state int

const (
	stateMenu state = iota
	stateInput
	stateMaze
)

type Model struct {
	maze     generator.Maze
	size     int
	width    int
	height   int
	state    state
	input    string
	viewport viewport.Model
}

func InitialModel(size int) Model {
	vp := viewport.New(
		viewport.WithWidth(80),
		viewport.WithHeight(20),
	)
	return Model{
		state:    stateMenu,
		viewport: vp,
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
		if m.state == stateMaze {
			m.updateMazeViewport()
		}
	case tea.KeyMsg:
		switch m.state {
		case stateMenu:
			switch msg.String() {
			case "p":
				m.state = stateInput
				m.input = ""
			case "q", "ctrl+c":
				return m, tea.Quit
			}

		case stateInput:
			switch msg.String() {
			case "enter":
				size, err := strconv.Atoi(m.input)
				if err == nil && size > 0 && size <= 50 {
					m.size = size
					m.maze = generator.MazeGenerator(size)
					m.updateMazeViewport()
					m.viewport.GotoTop()
					m.state = stateMaze
				}
			case "backspace":
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			case "escape":
				m.state = stateMenu
			case "q", "ctrl+c":
				return m, tea.Quit
			default:
				ch := msg.String()
				if len(ch) == 1 && ch[0] >= '0' && ch[0] <= '9' {
					m.input += ch
				}
			}
		case stateMaze:
			switch msg.String() {
			case "r":
				m.input = ""
				m.state = stateInput
			case "q", "ctrl+c":
				return m, tea.Quit
			}
			var cmd tea.Cmd
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}
	}
	return m, nil
}

func (m *Model) updateMazeViewport() {
	mazeText := render.RenderMazeWithLipgloss(m.maze)
	availableHeight := max(m.height-3, 1)
	mazeWidth := max(lipgloss.Width(mazeText), 1)
	mazeHeight := max(lipgloss.Height(mazeText), 1)

	m.viewport.SetWidth(min(m.width, mazeWidth))
	m.viewport.SetHeight(min(availableHeight, mazeHeight))
	m.viewport.SetContent(mazeText)
}

func (m Model) View() tea.View {
	var content, help string

	switch m.state {
	case stateMenu:
		content = "Terminal Maze Generator\n\nPress 'p' to play"
		help = "[p] play [q] quit"
	case stateInput:
		content = "Enter maze size (1-50):\n\n" + m.input + "█"
		help = "[enter] confirm  [esc] back  [q] quit"
	case stateMaze:
		content = m.viewport.View()
		help = "[up/down pgup/pgdn] scroll  [r] regenerate  [q] quit"
	}

	helpBlock := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, help)
	mainBlock := lipgloss.Place(m.width, m.height-3, lipgloss.Center, lipgloss.Center, content)
	s := lipgloss.JoinVertical(lipgloss.Left, mainBlock, "", helpBlock)
	return tea.NewView(s)
}
