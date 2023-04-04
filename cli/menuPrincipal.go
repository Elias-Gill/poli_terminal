package cli

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/styles"
)

func NewMainMenu() MainMenu {
	items := []list.Item{
        menuItem{Action: "horario", Tit: "Mi horario", Desc: "Revisa tu horario semanal y las fechas de examenes"},
		menuItem{Action: "calendario", Tit: "Calendario", Desc: "Mira en un calendario tus fechas de examenes"},
		menuItem{Action: "modHorario", Tit: "Modificar horario", Desc: "Realiza cambios al horario o crea uno desde 0"},
		menuItem{Action: "salir", Tit: "Salir", Desc: "Mas vale que sea para fiestear, ehemm, estudiar..."},
	}

	m := MainMenu{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}
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

type MainMenu struct {
	List     list.Model
	Selected bool
}

func (m MainMenu) Init() tea.Cmd {
	return nil
}

// actualizar el modelo
func (m MainMenu) Update(msg tea.Msg) (MainMenu, tea.Cmd) {
	exit_keys := map[string]struct{}{"q": {}, "esc": {}, "ctrl+c": {}}
    // handle special events
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			m.Selected = true
		}
		// si la tecla precionada es una de las de salir
		_, keyExit := exit_keys[msg.String()]
		if keyExit {
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
func (m MainMenu) View() string {
	return m.List.View()
}
