package cli

import (
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/excelParser"
)

func NewListaMats(height, width int, file string) (*ListaMats, error) {
	materias, err := excelParser.GetListaMaterias(file)
    if err != nil {
        return nil, err
    }
	items := []list.Item{}
	// cargar las materias disponibles
	for i, mat := range materias {
		items = append(items, itemLista{
			Tit:    mat.Nombre,
			Desc:   mat.Seccion + " - " + mat.Profesor,
			Action: strconv.Itoa(i)})
	}

	m := ListaMats{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.List.Title = "Lista de asignaturas"
	m.List.SetWidth(width)
	m.List.SetHeight(height)
	return &m, nil
}

type itemLista struct {
	Tit, Desc, Action string
}

func (i itemLista) Title() string       { return i.Tit }
func (i itemLista) Description() string { return i.Desc }
func (i itemLista) FilterValue() string { return i.Action }

type ListaMats struct {
	List     list.Model
	Selected bool
}

func (m ListaMats) Init() tea.Cmd {
	return nil
}

// actualizar el modelo
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
			return &m, tea.Quit
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return &m, cmd
}

// mostrar menu de seleccion
func (m ListaMats) View() string {
	return m.List.View()
}
