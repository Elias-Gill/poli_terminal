package armadorHorarios

import (
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

type selectMats struct {
	IsSelected bool // Determina si se selecciono una materia ("enter")
	list       list.Model
	materias   []*ep.Materia
	Focused    *ep.Materia
	Filtering  bool
	Quit       bool
	width      int
	height     int
}

// Retorna una nueva lista de materias. En caso de no poder abrirse el archivo excel, o este no ser valido,
// se retorna un error
func newSelectorMats() *selectMats {
	materias := cfm.GetUserConfig().MateriasExcel
	// Cargar las materias disponibles
	items := []list.Item{}
	for _, mat := range materias {
		items = append(items, itemLista{
			Tit:  mat.Nombre,
			Desc: mat.Seccion + " - " + mat.Profesor,
		})
	}

	// instanciar
	m := selectMats{
		list:       list.New(items, list.NewDefaultDelegate(), 0, 0),
		Quit:       false,
		materias:   materias,
		IsSelected: false,
	}
	m.list.Title = "Lista de asignaturas"
	return &m
}

func (m *selectMats) Update(msg tea.Msg) (tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	m.Filtering = m.list.FilterState().String() == "filtering"
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

	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
		m.height = msg.Height
		m.width = msg.Width
		return cmd
	}

	return nil
}

func (m *selectMats) View() string {
	// "meter" en una "caja" y centrar
	return lipgloss.PlaceHorizontal(m.width, lipgloss.Left, m.list.View())
}

// Retorna el elemento actualmente seleccionado
func (m *selectMats) indexOf() *ep.Materia {
	key := m.list.SelectedItem().FilterValue()
	for i, v := range m.materias {
		if v.Nombre == key {
			return m.materias[i]
		}
	}
	return nil
}
