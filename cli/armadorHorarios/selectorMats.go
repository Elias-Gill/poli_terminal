package armadorHorarios

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	ep "github.com/elias-gill/poli_terminal/excelParser"
	"github.com/elias-gill/poli_terminal/styles"
)

type itemLista struct {
	Tit, Desc string
}

func (i itemLista) Title() string       { return i.Tit }
func (i itemLista) Description() string { return i.Desc }
func (i itemLista) FilterValue() string { return i.Tit }

type SelectMats struct {
	list      list.Model
	materias  []ep.Materia
	Focus   ep.Materia
	Filtering bool
	Quit      bool
}

// WARN: cuidado con el camibo de paginas
//
// Retorna una nueva lista de materias. En caso de no poder abrirse el archivo excel, o este no ser valido,
// se retorna un error
func NewSelectorMats(f string) (SelectMats, error) {
	materias, err := ep.GetListaMaterias(f, 6)
	if err != nil {
		return SelectMats{}, err
	}
	// Cargar las materias disponibles
	items := []list.Item{}
	for i, mat := range materias {
		aux := itemLista{
			Tit:  mat.Nombre,
			Desc: mat.Seccion + " - " + mat.Profesor,
		}
		items = append(items, aux)
		materias[i].Nombre = aux.FilterValue()
	}

	// instanciar
	m := SelectMats{
		list:     list.New(items, list.NewDefaultDelegate(), 0, 0),
		Quit:     false,
		materias: materias,
	}
	m.list.Title = "Lista de asignaturas"
	m.list.SelectedItem()

	return m, nil
}

func (m SelectMats) Init() tea.Cmd { return nil }

func (m SelectMats) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// handle special events
	filtering := m.list.FilterState().String() == "filtering"
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		w, h := styles.DocStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-w)
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	if !filtering {
		m.Focus = m.materias[m.indexOf(m.list.SelectedItem().FilterValue())]
	}
	return m, cmd
}

func (m SelectMats) View() string {
	return m.list.View()
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
