package horario

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	ep "github.com/elias-gill/poli_terminal/excelParser"
)

type DisplayHorario struct {
	table table.Model
	Quit  bool
}

func (m DisplayHorario) Init() tea.Cmd { return nil }

func (m DisplayHorario) Update(msg tea.Msg) (DisplayHorario, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// si la tecla precionada es una de las de salir
		keyExit := msg.String() == "q" || msg.String() == "esc"
		if keyExit {
			m.Quit = true
			return m, nil
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m DisplayHorario) View() string {
	var baseStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("13"))

	return baseStyle.Render(m.table.View()) + "\n\n" + 
    baseStyle.Render()
}

func NewHorario(m []ep.Materia) DisplayHorario {
	columns := []table.Column{
		{Title: "Asignatura", Width: 18},
		{Title: "Profesor", Width: 18},
		{Title: "Semestre", Width: 8},
		{Title: "Seccion", Width: 7},
		{Title: "Parcial 1", Width: 18},
		{Title: "Parcial 2", Width: 18},
		{Title: "Final 1", Width: 18},
		{Title: "Final 2", Width: 18},
	}

	rows := []table.Row{}
	for _, v := range m {
		rows = append(rows, table.Row{
			v.Nombre,
			v.Profesor,
			strconv.Itoa(v.Semestre),
			v.Seccion,
			v.Parcial1,
			v.Parcial2,
			v.Final1,
			v.Final2,
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(8),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return DisplayHorario{
		table: t,
		Quit:  false,
	}
}
