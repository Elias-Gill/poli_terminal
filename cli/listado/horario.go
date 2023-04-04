package listado

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elias-gill/poli_terminal/excelParser"
	ep "github.com/elias-gill/poli_terminal/excelParser"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type Horario struct {
	table    table.Model
	Quit     bool
	infoMode bool
    infoComp infoMateria
}

func (m Horario) Init() tea.Cmd { return nil }

func (m Horario) Update(msg tea.Msg) (Horario, tea.Cmd) {
	var cmd tea.Cmd
	options := map[string]struct{}{"q": {}, "esc": {}}

    // modo de informacion de la materia
    if m.infoMode {
        m.infoComp, cmd = m.infoComp.Update(msg)
        if m.infoComp.Quit {
            m.infoMode = false
        }
        return m, cmd
    }

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			m.infoMode = true
            m.infoComp = newInfoMateria(excelParser.Materia{})
            return m, nil
		}

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

func (m Horario) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func NewHorario(m []ep.Materia) Horario {
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

	return Horario{
		table: t,
		Quit:  false,
	}
}
