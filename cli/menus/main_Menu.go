package menus

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/cli/constants"
	"github.com/elias-gill/poli_terminal/styles"
)

// struct para los items del menu
type menuItem struct {
	Tit, Desc, Action string
}

func (i menuItem) Title() string       { return i.Tit }
func (i menuItem) Description() string { return i.Desc }
func (i menuItem) FilterValue() string { return i.Action }

// struct principal
type MainMenu struct {
	List       list.Model
	IsSelected bool
}

const (
	horario       = "horario"
	calendario    = "calendario"
	scheduleMaker = "scheduleMaker"
	configMenu    = "configMenu"
	salir         = "salir"
)

// generates a new instance of the main menu
func NewMainMenu() MainMenu {
	items := []list.Item{
		menuItem{Action: "horario", Tit: "Mi horario", Desc: "Revisa tu horario semanal y las fechas de examenes"},
		menuItem{Action: "calendario", Tit: "Calendario", Desc: "Mira en un calendario tus fechas de examenes"},
		menuItem{Action: "scheduleMaker", Tit: "Modificar horario", Desc: "Realiza cambios en el horario (primero debes configurar el excel en 'Configuraciones')"},
		menuItem{Action: "configMenu", Tit: "Configuracion", Desc: "Cambia las configuraciones del sistema"},
		menuItem{Action: "salir", Tit: "Salir", Desc: "Mas vale que sea para fiestear, ehemm, estudiar..."},
	}

	m := MainMenu{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.List.Title = "Mi Politerminal"
	m.List.SetFilteringEnabled(false)
	return m
}

func (m MainMenu) Init() tea.Cmd {
	return nil
}

// actualizar el modelo
func (m MainMenu) Update(msg tea.Msg) (constants.Component, tea.Cmd) {
	// handle special events
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, m.changeMode()
		case "esc", "q":
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
func (m MainMenu) Render() string {
	return m.List.View()
}

func (m MainMenu) changeMode() tea.Cmd {
	var cmd tea.Cmd = nil
	switch m.List.SelectedItem().FilterValue() {
	case horario:
		constants.CurrentMode = constants.InScheduleDisplayer

	case calendario:
		constants.CurrentMode = constants.InCalendar

	case salir:
		cmd = tea.Quit

	case scheduleMaker:
		constants.CurrentMode = constants.InScheduleMaker

	case configMenu:
		constants.CurrentMode = constants.InConfigMenu
	}
	return cmd
}
