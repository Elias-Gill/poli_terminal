package armadorHorarios

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	ep "github.com/elias-gill/poli_terminal/excelParser"
	"github.com/elias-gill/poli_terminal/styles"
)

type infoMateria struct {
	Quit    bool
	materia ep.Materia
	width   int
	height  int
}

func newInfoMateria(m ep.Materia) infoMateria {
	return infoMateria{materia: m, Quit: false}
}

func (i infoMateria) Init() tea.Cmd { return nil }

func (i infoMateria) Update(msg tea.Msg) (infoMateria, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		i.width = msg.Width
		i.height = msg.Height
	}
	return i, nil
}

func (i infoMateria) View() string {
	res := styles.DoneStyle.Render(i.materia.Profesor) +
		"\n" + styles.GoodStyle.Render("Parciales: \t") +
		"\n" + i.materia.Parcial1 +
		"\n" + i.materia.Parcial2 +
		"\n" + styles.GoodStyle.Render("Finales: \t") +
		"\n" + i.materia.Final1 +
		"\n" + i.materia.Final2 + "\n"

	return lipgloss.PlaceHorizontal(i.width, lipgloss.Left, res)
}

func (i infoMateria) ChangeMateria(m ep.Materia) infoMateria {
	i.materia = m
	return i
}
