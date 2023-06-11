package schedule

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elias-gill/poli_terminal/configManager"
	ep "github.com/elias-gill/poli_terminal/excelParser"
)

type ScheduleDisplayer struct {
	tablaMats table.Model
	tablaDias table.Model
	Quit      bool
}

func (m ScheduleDisplayer) Init() tea.Cmd { return nil }

func (m ScheduleDisplayer) Update(msg tea.Msg) (ScheduleDisplayer, tea.Cmd) {
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

	m.tablaMats, cmd = m.tablaMats.Update(msg)
	return m, cmd
}

func (m ScheduleDisplayer) View() string {
	var baseStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("23"))

	return baseStyle.Render(m.tablaDias.View()) + "\n\n" +
		baseStyle.Render(m.tablaMats.View())
}

func NewScheduleDisplayer() ScheduleDisplayer {
	m := configManager.GetUserConfig().MateriasUsuario
	return ScheduleDisplayer{
		tablaMats: nuevaTablaMats(m),
		tablaDias: nuevaTablaDias(m),
		Quit:      false,
	}
}

func nuevaTablaDias(m []*ep.Materia) table.Model {
	columns := []table.Column{
		{Title: "", Width: 18},
		{Title: "Lunes", Width: 18},
		{Title: "Martes", Width: 18},
		{Title: "Miercoles", Width: 18},
		{Title: "Jueves", Width: 18},
		{Title: "Viernes", Width: 18},
		{Title: "Sabado", Width: 18},
	}

	rows := []table.Row{}
	for _, v := range m {
		rows = append(rows, table.Row{
			v.Nombre,
			v.Dias.Lunes,
			v.Dias.Martes,
			v.Dias.Miercoles,
			v.Dias.Jueves,
			v.Dias.Viernes,
			v.Dias.Sabado,
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(8),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.Bold(false)

	t.SetStyles(s)
	return t
}

func nuevaTablaMats(m []*ep.Materia) table.Model {
	columns := []table.Column{
		{Title: "Asignatura", Width: 18},
		{Title: "Profesor", Width: 18},
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
		table.WithFocused(false),
		table.WithHeight(8),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.Bold(false)

	return t
}
