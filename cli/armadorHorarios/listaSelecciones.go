package armadorHorarios

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	ep "github.com/elias-gill/poli_terminal/excelParser"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type listSelecs struct {
	table  table.Model
	lista  []ep.Materia
	height int
	width  int
	Quit   bool
}

func (m listSelecs) Init() tea.Cmd { return nil }

func (m listSelecs) Update(msg tea.Msg) (listSelecs, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "x" {
			if len(m.table.Rows()) > 0 {
				m.table, cmd = m.table.Update(msg)
				var i int
				m.lista, i = m.DelMateria()
				m.table.SetRows(m.nuevasFilas())
				m.table.SetCursor(i)
				return m, cmd
			}
		}
	case tea.WindowSizeMsg:
		x, y := baseStyle.GetFrameSize()
		m.table.SetWidth(msg.Width - x)
		m.table.SetHeight(msg.Height - y)
		return m, cmd
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m listSelecs) View() string {
	var style = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63"))
	return style.Render(lipgloss.PlaceHorizontal(m.width, lipgloss.Right, m.table.View()))
}

// retorna una nueva lista de materias
func newLista(m []ep.Materia) listSelecs {
	return listSelecs{
		table: construirTabla(m),
		Quit:  false,
	}
}

func construirTabla(m []ep.Materia) table.Model {
	columns := []table.Column{
		{Title: "Asignatura", Width: 30},
		{Title: "Seccion", Width: 7},
	}

	rows := []table.Row{}
	for _, v := range m {
		rows = append(rows, table.Row{
			v.Nombre,
			v.Seccion,
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
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

	return t
}

// Anade la nueva materia proporcionada a la lista
func (l listSelecs) AddMateria(mat ep.Materia) listSelecs {
	// buscar que no se repita
	for _, v := range l.lista {
		if v.Nombre == mat.Nombre {
			return l
		}
	}
	l.lista = append(l.lista, mat)
	l.table.SetRows(l.nuevasFilas())
	return l
}

// Elimina de la lista la materias actualmente enfocada.
//
// Retorna una nueva lista y el indice donde se debe colocar de nuevo el foco
// de la lista
func (l listSelecs) DelMateria() ([]ep.Materia, int) {
	selec := l.table.SelectedRow()[0]
	var aux []ep.Materia
	index := 1
	for i, m := range l.lista {
		if m.Nombre != selec {
			aux = append(aux, m)
			continue
		}
		index = i
	}
	return aux, index - 1
}

func (l listSelecs) nuevasFilas() []table.Row {
	rows := []table.Row{}
	for _, v := range l.lista {
		rows = append(rows, table.Row{
			v.Nombre,
			v.Seccion,
		})
	}
	return rows
}
