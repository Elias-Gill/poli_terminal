package cli

import (
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/excelParser"
)

// Retorna una nueva lista de materias. En caso de no poder abrirse el archivo excel, o este no ser valido,
// se retorna un error
func NewListaMats(height, width int, file string) (*ListaMats, error) {
	materias, err := excelParser.GetListaMaterias(file)
	if err != nil {
		return nil, err
	}
	items := []list.Item{}
	// Cargar las materias disponibles
	for i, mat := range materias {
		items = append(items, itemLista{
			Tit:  mat.Nombre + " #" + strconv.Itoa(i),
			Desc: mat.Seccion + " - " + mat.Profesor,
		})
	}

	m := ListaMats{
		List:     list.New(items, list.NewDefaultDelegate(), 0, 0),
		Selected: false,
		Quit:     false,
	}
	m.List.Title = "Lista de asignaturas"
	h, v := docStyle.GetFrameSize()
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
	Selected bool
	Quit     bool
}

func (m ListaMats) Init() tea.Cmd {
	return nil
}

func (m ListaMats) Update(msg tea.Msg) (*ListaMats, tea.Cmd) {
	options := map[string]struct{}{"q": {}, "esc": {}, "ctrl+c": {}}
	// handle special events
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			m.Selected = true
		}
		// si la tecla precionada es una de las de salir
		_, keyExit := options[msg.String()]
		if keyExit {
            m.Quit = true
			return &m, nil
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return &m, cmd
}

func (m ListaMats) View() string {
	return m.List.View()
}
