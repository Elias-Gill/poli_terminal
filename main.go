package main

import (
    // "github.com/elias-gill/poli_terminal/excelParser"
	"fmt"
	"os"

    "github.com/elias-gill/poli_terminal/cli"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	m := cli.NewApp()
    // run
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
