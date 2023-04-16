package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/cli"
	"github.com/elias-gill/poli_terminal/configManager"
	"github.com/elias-gill/poli_terminal/excelParser"
)

func main() {
	m := cli.NewApp()
	// run
	p := tea.NewProgram(m, tea.WithAltScreen())
    // precargar el excel
	go excelParser.OpenExcelFile(configManager.GetUserConfig().FHorario)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		excelParser.CloseExcel()
		os.Exit(1)
	}
}
