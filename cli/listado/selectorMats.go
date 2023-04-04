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
	infoMode bool
	focus    int
	infoMat  infoMateria
	Quit     bool
}

// Retorna una nueva lista de materias. En caso de no poder abrirse el archivo excel, o este no ser valido,
// se retorna un error
func NewSelectorMats(height, width int, file string) (*SelectMats, error) {
	// WARN: cuidado con el camibo de paginas
	materias, err := excelParser.GetListaMaterias(file, 6)
	if err != nil {
		return nil, err
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
		infoMode: false,
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

func (m SelectMats) Update(msg tea.Msg) (*SelectMats, tea.Cmd) {
	if m.infoMode {
		var cmd tea.Cmd
		m.infoMat, cmd = m.infoMat.Update(msg)
		if m.infoMat.Quit {
			m.infoMode = false
		}
		return &m, cmd
	}

	options := map[string]struct{}{"q": {}, "esc": {}}
	// handle special events
	switch msg := msg.(type) {
	case tea.KeyMsg:
		filtering := m.List.FilterState().String() == "filtering"
		// info de materia
		if msg.String() == "i" && !filtering {
			i := m.indexOf(m.List.SelectedItem().FilterValue())
			m.infoMat = newInfoMateria(m.materias[i])
			m.infoMode = true
			return &m, nil
		}

		// TODO: cuadros de info de materias

		// si la tecla precionada es una de las de salir
		_, keyExit := options[msg.String()]
		if keyExit && !filtering {
			m.Quit = true
			return &m, nil
		}

	case tea.WindowSizeMsg:
		h, v := styles.DocStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v-6)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return &m, cmd
}

func (m SelectMats) View() string {
	if m.infoMode {
		return m.infoMat.View()
	}
	return m.List.View() + "\n\n"
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
