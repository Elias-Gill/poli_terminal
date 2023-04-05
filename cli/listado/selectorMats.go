package listado

import (
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/excelParser"
	"github.com/elias-gill/poli_terminal/styles"
)

type itemLista struct {
	Tit, Desc string
}

func (i itemLista) Title() string       { return i.Tit }
func (i itemLista) Description() string { return i.Desc }
func (i itemLista) FilterValue() string { return i.Tit }

type SelectMats struct {
	List     list.Model
	materias []excelParser.Materia
	focused  excelParser.Materia
	Quit     bool
}

// WARN: cuidado con el camibo de paginas
//
// Retorna una nueva lista de materias. En caso de no poder abrirse el archivo excel, o este no ser valido,
// se retorna un error
func NewSelectorMats(f string) (SelectMats, error) {
	materias, err := excelParser.GetListaMaterias(f, 6)
	if err != nil {
		return SelectMats{}, err
	}
	// Cargar las materias disponibles
	items := []list.Item{}
	for i, mat := range materias {
		aux := itemLista{
			Tit:  "#" + strconv.Itoa(i) + "  " + mat.Nombre,
			Desc: mat.Seccion + " - " + mat.Profesor,
		}
		items = append(items, aux)
		materias[i].Nombre = aux.FilterValue()
	}

	// instanciar
	m := SelectMats{
		List:     list.New(items, list.NewDefaultDelegate(), 0, 0),
		Quit:     false,
		materias: materias,
	}
	m.List.Title = "Lista de asignaturas"
	m.List.SelectedItem()

	return m, nil
}

func (m SelectMats) Init() tea.Cmd { return nil }

func (m SelectMats) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	options := map[string]struct{}{"q": {}, "esc": {}}

	// handle special events
	filtering := m.List.FilterState().String() == "filtering"
	switch msg := msg.(type) {
	case tea.KeyMsg:

		// si la tecla precionada es una de las de salir
		_, keyExit := options[msg.String()]
		if keyExit && !filtering {
			m.Quit = true
			return m, nil
		}

	case tea.WindowSizeMsg:
		h, v := styles.DocStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v-6)
	}
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	if !filtering {
		m.focused = m.materias[m.indexOf(m.List.SelectedItem().FilterValue())]
	}
	return m, cmd
}

func (m SelectMats) View() string {
	return m.List.View()
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
