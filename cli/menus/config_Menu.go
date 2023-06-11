package menus

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/styles"
)

func NewConfigMenu() ConfigMenu {
	items := []list.Item{
		menuItem{Action: "Excel", Tit: "Excel", Desc: "Cambia el archivo excel que se lee"},
	}

	m := ConfigMenu{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.List.Title = "Settings"
	m.List.SetFilteringEnabled(false)
	return m
}

type menuConfigItem struct {
	Tit, Desc, Action string
}

func (i menuConfigItem) Title() string       { return i.Tit }
func (i menuConfigItem) Description() string { return i.Desc }
func (i menuConfigItem) FilterValue() string { return i.Action }

type ConfigMenu struct {
	Quit     bool
	List     list.Model
	Selected bool
}

func (m ConfigMenu) Init() tea.Cmd {
	return nil
}

// actualizar el modelo
func (m ConfigMenu) Update(msg tea.Msg) (ConfigMenu, tea.Cmd) {
	// handle special events
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			m.Selected = true
		}
		// si la tecla precionada es una de las de salir
		if msg.String() == "q" || msg.String() == "esc" {
            m.Quit = true
			return m, nil
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
func (m ConfigMenu) View() string {
	return m.List.View()
}
