package model

import (
	"strconv"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
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
	stateWin
)

type menuKeyMap struct{}

func (k menuKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		key.NewBinding(key.WithKeys("p"), key.WithHelp("p", "play")),
		key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "quit")),
	}
}

func (k menuKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

type inputKeyMap struct{}

func (k inputKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "confirm")),
		key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "back")),
		key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "quit")),
	}
}

func (k inputKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

type mazeKeyMap struct{}

func (k mazeKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		key.NewBinding(key.WithKeys("up", "down", "left", "right"), key.WithHelp("↑/↓/←/→", "move")),
		key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "regenerate")),
		key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "quit")),
	}
}

func (k mazeKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

type winKeyMap struct{}

func (k winKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "play again")),
		key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "quit")),
	}
}

func (k winKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

type Model struct {
	maze      generator.Maze
	size      int
	width     int
	height    int
	state     state
	input     string
	viewport  viewport.Model
	help      help.Model
	playerRow int
	playerCol int
}

func InitialModel() Model {
	vp := viewport.New(
		viewport.WithWidth(80),
		viewport.WithHeight(20),
	)
	return Model{
		state:    stateMenu,
		viewport: vp,
		help:     help.New(),
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
					m.playerRow = m.maze.Start[0]
					m.playerCol = m.maze.Start[1]
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
			case "up":
				if !m.maze.Grid[m.playerRow][m.playerCol].Up {
					m.playerRow--
					m.updateMazeViewport()
					if m.playerRow == m.maze.End[0] && m.playerCol == m.maze.End[1] {
						m.state = stateWin
					}
				}
				return m, nil
			case "down":
				if !m.maze.Grid[m.playerRow][m.playerCol].Down {
					m.playerRow++
					m.updateMazeViewport()
					if m.playerRow == m.maze.End[0] && m.playerCol == m.maze.End[1] {
						m.state = stateWin
					}
				}
				return m, nil
			case "left":
				if !m.maze.Grid[m.playerRow][m.playerCol].Left {
					m.playerCol--
					m.updateMazeViewport()
					if m.playerRow == m.maze.End[0] && m.playerCol == m.maze.End[1] {
						m.state = stateWin
					}
				}
				return m, nil
			case "right":
				if !m.maze.Grid[m.playerRow][m.playerCol].Right {
					m.playerCol++
					m.updateMazeViewport()
					if m.playerRow == m.maze.End[0] && m.playerCol == m.maze.End[1] {
						m.state = stateWin
					}
				}
				return m, nil
			case "r":
				m.input = ""
				m.state = stateInput
				return m, nil
			case "q", "ctrl+c":
				return m, tea.Quit
			}
			var cmd tea.Cmd
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		case stateWin:
			switch msg.String() {
			case "r":
				m.input = ""
				m.state = stateInput
			case "q", "ctrl+c":
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m *Model) updateMazeViewport() {
	mazeText := render.RenderMazeWithLipgloss(m.maze, m.playerRow, m.playerCol)
	availableHeight := max(m.height-3, 1)

	m.viewport.SetWidth(m.width)
	m.viewport.SetHeight(availableHeight)
	centered := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, mazeText)
	m.viewport.SetContent(centered)

	playerLine := 2*m.playerRow + 1
	margin := availableHeight / 4
	topThreshold := m.viewport.YOffset() + margin
	bottomThreshold := m.viewport.YOffset() + availableHeight - margin - 1

	if playerLine < topThreshold {
		m.viewport.SetYOffset(max(playerLine-margin, 0))
	} else if playerLine > bottomThreshold {
		m.viewport.SetYOffset(playerLine - availableHeight + margin + 1)
	}
}

func (m Model) View() tea.View {
	var content string
	var km help.KeyMap

	switch m.state {
	case stateMenu:
		content = "Terminal Maze Generator\n\nPress 'p' to play"
		km = menuKeyMap{}
	case stateInput:
		content = "Enter maze size (1-50):\n\n" + m.input + "█"
		km = inputKeyMap{}
	case stateMaze:
		content = m.viewport.View()
		km = mazeKeyMap{}
	case stateWin:
		content = "Congratulations! You solved the maze!\n\nPress 'r' to play again"
		km = winKeyMap{}
	}

	m.help.SetWidth(m.width)
	helpView := m.help.View(km)
	helpBlock := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, helpView)
	mainBlock := lipgloss.Place(m.width, m.height-3, lipgloss.Center, lipgloss.Center, content)
	s := lipgloss.JoinVertical(lipgloss.Left, mainBlock, "", helpBlock)
	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
