package schedule

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	cfm "github.com/elias-gill/poli_terminal/configManager"
	ep "github.com/elias-gill/poli_terminal/excelParser"
)

type itemLista struct {
	Tit, Desc string
}

func (i itemLista) Title() string       { return i.Tit }
func (i itemLista) Description() string { return i.Desc }
func (i itemLista) FilterValue() string { return i.Tit }

type selector struct {
	list   list.Model
	width  int
	height int
	// estados
	IsSelected bool // Determina si se selecciono una materia ("enter")
	Filtering  bool
	Quit       bool
	// materias
	materias []*ep.Materia
	Focused  *ep.Materia
}

// Retorna una nueva lista de materias. En caso de no poder abrirse el archivo excel, o este no ser valido,
// se retorna un error
func newSelectorMats() (*selector, error) {
	materias := cfm.GetUserConfig().MateriasExcel
    if materias == nil {
        return nil, fmt.Errorf("Error al parsear requerir la config del usuario")
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
	m := selector{
		list:       list.New(items, list.NewDefaultDelegate(), 0, 0),
		Quit:       false,
		materias:   materias,
		Focused:    &ep.Materia{},
		IsSelected: false,
	}
	m.list.Title = "Lista de asignaturas"
	return &m, nil
}

func (m *selector) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	if !m.Filtering {
		m.Focused = m.indexOf()
	}

	// handle special events
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" && !m.Filtering {
			if m.Focused != nil {
				m.IsSelected = true
			}
			return cmd
		}
        if msg.String() == "q" || msg.String() == tea.KeyEsc.String() {
            if !m.Filtering {
                m.Quit = true
                return nil
            }
        }

	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
		m.height = msg.Height
		m.width = msg.Width
		return cmd
	}

	m.Filtering = m.list.FilterState().String() == "filtering"
	return cmd
}

func (m *selector) View() string {
	// "meter" en una "caja" y centrar
	return lipgloss.PlaceHorizontal(m.width, lipgloss.Left, m.list.View())
}

// Retorna el elemento actualmente seleccionado
func (m *selector) indexOf() *ep.Materia {
	key := m.list.SelectedItem().FilterValue()
	for i, v := range m.materias {
		if v.Nombre == key {
			return m.materias[i]
		}
	}
	return &ep.Materia{}
}
