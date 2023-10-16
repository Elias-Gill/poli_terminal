package menus

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/cli/constants"
	"github.com/elias-gill/poli_terminal/cli/prompts"
	"github.com/elias-gill/poli_terminal/styles"
)

type menuConfigItem struct {
	Tit, Desc, Action string
}

type ConfigMenu struct {
	mode int
	List list.Model
	// components
	fileTree *prompts.Filetree
}

const (
    inMenu = iota
	inFileTree
)

func NewConfigMenu() ConfigMenu {
	items := []list.Item{
		menuItem{Action: "Excel", Tit: "Excel", Desc: "Cambia el archivo excel que se lee"},
	}

    m := ConfigMenu{List: list.New(items, list.NewDefaultDelegate(), 0, 0), mode: inMenu}
	m.List.Title = "Settings"
	m.List.SetFilteringEnabled(false)
	return m
}

func (i menuConfigItem) Title() string       { return i.Tit }
func (i menuConfigItem) Description() string { return i.Desc }
func (i menuConfigItem) FilterValue() string { return i.Action }

func (m ConfigMenu) Init() tea.Cmd {
	return nil
}

func (m ConfigMenu) Update(msg tea.Msg) (constants.Component, tea.Cmd) {
	if m.mode == inFileTree {
		var cmd tea.Cmd
		cmd = m.fileTree.Update(msg)
        if m.fileTree.Quit {
            m.mode = inMenu
        }
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m.changeMode()

		case "q", "esc":
            constants.CurrentMode = constants.InMainMenu
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
func (m ConfigMenu) Render() string {
    if m.mode == inFileTree {
        return m.fileTree.View()
    }
	return m.List.View()
}

func (m ConfigMenu) changeMode() (ConfigMenu, tea.Cmd) {
	switch m.List.SelectedItem().FilterValue() {
	case "Excel":
		m.mode = inFileTree
		m.fileTree = prompts.NewFiletree()
		m.fileTree.Update(
			tea.WindowSizeMsg{
				Width:  m.List.Width(),
				Height: m.List.Height(),
			},
		)
	}
	return m, nil
}
