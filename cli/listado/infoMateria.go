package listado

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	ep "github.com/elias-gill/poli_terminal/excelParser"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type infoMateria struct {
	table table.Model
	Quit  bool
}

func (m infoMateria) Init() tea.Cmd { return nil }

func (m infoMateria) Update(msg tea.Msg) (infoMateria, tea.Cmd) {
	var cmd tea.Cmd
	options := map[string]struct{}{"q": {}, "esc": {}, "ctrl+c": {}}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// si la tecla precionada es una de las de salir
		_, keyExit := options[msg.String()]
		if keyExit {
			m.Quit = true
			return m, nil
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m infoMateria) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func NewInfoMateria(m []ep.Materia) infoMateria {
	columns := []table.Column{
		{Title: "Asignatura", Width: 25},
		{Title: "Profesor", Width: 32},
		{Title: "Semestre", Width: 10},
		{Title: "Seccion", Width: 10},
		{Title: "Parcial 1", Width: 21},
		{Title: "Parcial 2", Width: 21},
		{Title: "Final 1", Width: 21},
		{Title: "Final 2", Width: 21},
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
		table.WithHeight(7),
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

	return infoMateria{
		table: t,
		Quit:  false,
	}
}
