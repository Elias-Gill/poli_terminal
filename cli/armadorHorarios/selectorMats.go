package armadorHorarios

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	ep "github.com/elias-gill/poli_terminal/excelParser"
)

type itemLista struct {
	Tit, Desc string
}

func (i itemLista) Title() string       { return i.Tit }
func (i itemLista) Description() string { return i.Desc }
func (i itemLista) FilterValue() string { return i.Tit }

type SelectMats struct {
	Selected  bool
	list      list.Model
	materias  []ep.Materia
	Focused   ep.Materia
	Filtering bool
	Quit      bool
	width     int
	height    int
}

// WARN: cuidado con el camibo de paginas
//
// Retorna una nueva lista de materias. En caso de no poder abrirse el archivo excel, o este no ser valido,
// se retorna un error
func newSelectorMats(f string) SelectMats {
	materias, err := ep.GetListaMaterias(f, 6)
	if err != nil {
		panic(err)
	}
	// Cargar las materias disponibles
	items := []list.Item{}
	for _, mat := range materias {
		items = append(items, itemLista{
			Tit:  mat.Nombre,
			Desc: mat.Seccion + " - " + mat.Profesor,
		})
	}

	// instanciar
	m := SelectMats{
		list:     list.New(items, list.NewDefaultDelegate(), 0, 0),
		Quit:     false,
		materias: materias,
		Selected: false,
	}
	m.list.Title = "Lista de asignaturas"

	return m
}

func (m SelectMats) Init() tea.Cmd { return nil }

func (m SelectMats) Update(msg tea.Msg) (SelectMats, tea.Cmd) {
	var cmd tea.Cmd
	// handle special events
	switch msg := msg.(type) {
	case tea.KeyMsg:
		f := m.list.FilterState().String() == "filtering"
		if msg.String() == "enter" && !f {
			m.Selected = true
			return m, cmd
		}

	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
		m.height = msg.Height
		m.width = msg.Width
		return m, cmd
	}

	m.list, cmd = m.list.Update(msg)
	m.Filtering = m.list.FilterState().String() == "filtering"
	if !m.Filtering {
		m.Focused = m.materias[m.indexOf(m.list.SelectedItem().FilterValue())]
	}
	return m, nil
}

func (m SelectMats) View() string {
	return lipgloss.PlaceHorizontal(m.width, lipgloss.Left, m.list.View())
}

// buscar el valor seleccionado en la lista
func (m SelectMats) indexOf(key string) int {
	for i, v := range m.materias {
		if v.Nombre == key {
			return i
		}
	}
	return 0
}
