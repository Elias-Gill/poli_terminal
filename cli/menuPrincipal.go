package cli

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/styles"
)

// generates a new instance of the main menu
func NewMainMenu() MenuPrincipal {
	items := []list.Item{
		menuItem{Action: "horario", Tit: "Mi horario", Desc: "Revisa tu horario semanal y las fechas de examenes"},
		menuItem{Action: "calendario", Tit: "Calendario", Desc: "Mira en un calendario tus fechas de examenes"},
		menuItem{Action: "scheduleMaker", Tit: "Modificar horario", Desc: "Realiza cambios en el horario (primero debes configurar el excel en 'Configuraciones')"},
		menuItem{Action: "configMenu", Tit: "Configuracion", Desc: "Cambia las configuraciones del sistema"},
		menuItem{Action: "salir", Tit: "Salir", Desc: "Mas vale que sea para fiestear, ehemm, estudiar..."},
	}

	m := MenuPrincipal{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.List.Title = "Mi Politerminal"
	m.List.SetFilteringEnabled(false)
	return m
}

type menuItem struct {
	Tit, Desc, Action string
}

func (i menuItem) Title() string       { return i.Tit }
func (i menuItem) Description() string { return i.Desc }
func (i menuItem) FilterValue() string { return i.Action }

type MenuPrincipal struct {
	List       list.Model
	IsSelected bool
}

func (m MenuPrincipal) Init() tea.Cmd {
	return nil
}

// actualizar el modelo
func (m MenuPrincipal) Update(msg tea.Msg) (MenuPrincipal, tea.Cmd) {
	// handle special events
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			m.IsSelected = true
		}
		// si la tecla precionada es una de las de salir
		if msg.String() == "q" || msg.String() == "esc" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		h, v := styles.DocStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

// mostrar menu de seleccion
func (m MenuPrincipal) View() string {
	return m.List.View()
}
