package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/ChamikaUluwatta/TerminalMaze/model"
)

func main() {
	p := tea.NewProgram(model.InitialModel(4))
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
