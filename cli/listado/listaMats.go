package listado

import (
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/excelParser"
	"github.com/elias-gill/poli_terminal/styles"
)

// Retorna una nueva lista de materias. En caso de no poder abrirse el archivo excel, o este no ser valido,
// se retorna un error
func NewListaMats(height, width int, file string) (*ListaMats, error) {
	// WARN: cuidado con el camibo de paginas
	materias, err := excelParser.GetListaMaterias(file, 6)
	if err != nil {
		return nil, err
	}
	items := []list.Item{}
	// Cargar las materias disponibles
	for i, mat := range materias {
		aux := itemLista{
			Tit:  mat.Nombre + " #" + strconv.Itoa(i),
			Desc: mat.Seccion + " - " + mat.Profesor,
		}
		items = append(items, aux)
		materias[i].Nombre = aux.FilterValue()
	}

	m := ListaMats{
		List:     list.New(items, list.NewDefaultDelegate(), 0, 0),
		Selected: false,
		Quit:     false,
		materias: materias,
	}
	m.List.Title = "Lista de asignaturas"
	h, v := styles.DocStyle.GetFrameSize()
	m.List.SetWidth(width - v)
	m.List.SetHeight(height - h)
	m.List.SelectedItem()

	return &m, nil
}

type itemLista struct {
	Tit, Desc string
}

func (i itemLista) Title() string       { return i.Tit }
func (i itemLista) Description() string { return i.Desc }
func (i itemLista) FilterValue() string { return i.Tit }

type ListaMats struct {
	List     list.Model
	materias []excelParser.Materia
	Selected bool
	infoMat  infoMateria
	Quit     bool
}

func (m ListaMats) Init() tea.Cmd {
	return nil
}

func (m ListaMats) Update(msg tea.Msg) (*ListaMats, tea.Cmd) {
	if m.Selected {
		var cmd tea.Cmd
		m.infoMat, cmd = m.infoMat.Update(msg)
		if m.infoMat.Quit {
			m.Selected = false
		}
		return &m, cmd
	}

	options := map[string]struct{}{"q": {}, "esc": {}, "ctrl+c": {}}
	// handle special events
	switch msg := msg.(type) {
	case tea.KeyMsg:
		filtering := m.List.FilterState().String() == "filtering"
		if msg.String() == "enter" && !filtering {
			i := m.indexOf(m.List.SelectedItem().FilterValue())
			m.infoMat = NewInfoMateria([]excelParser.Materia{m.materias[i]})
			m.Selected = true
			return &m, nil
		}
		// si la tecla precionada es una de las de salir
		_, keyExit := options[msg.String()]
		if keyExit && !filtering{
			m.Quit = true
			return &m, nil
		}

	case tea.WindowSizeMsg:
		h, v := styles.DocStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return &m, cmd
}

func (m ListaMats) View() string {
	if m.Selected {
		return m.infoMat.View()
	}
	return m.List.View()
}

// buscar el valor seleccionado en la lista
func (m ListaMats) indexOf(key string) int {
	for i, v := range m.materias {
		if v.Nombre == key {
			return i
		}
	}
	return 0
}
