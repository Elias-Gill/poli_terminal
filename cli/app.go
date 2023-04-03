package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/configManager"
)

// modos
const (
	inMenu = iota
	inCalendar
	inHorario
	inSelection
	inListMats
	inAlert
)

type App struct {
	Mode      int
	appWith   int
	appHeight int
	config    configManager.Configurations
	// components
	mainMenu  MainMenu
	listaMats *ListaMats
}

func NewApp() App {
	return App{
		mainMenu: NewMainMenu(),
		config: configManager.GetUserConfig(),
		Mode:     inMenu,
	}
}

func (m App) Init() tea.Cmd {
	return nil
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// actualizar el tamano de la app
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.appWith = msg.Width
		a.appHeight = msg.Height
	}

	// handle events
	switch a.Mode {
	case inCalendar: // TODO: implementar
		return a, cmd

	case inAlert: // TODO: implementar
		return a, cmd

	case inListMats:
		a.listaMats, cmd = a.listaMats.Update(msg)
		return a, cmd
	}

	// por defecto nos encontramos en el menu principal
	a.mainMenu, cmd = a.mainMenu.Update(msg)
	if a.mainMenu.Selected {
		return a.selectMode()
	}
	return a, cmd
}

// selecciona la vista dependiendo del estado de la aplicacion
func (m App) View() string {
	switch m.Mode {
	case inListMats:
		return docStyle.Render(m.listaMats.View())

	case inHorario:
		return "Sorry, this is not implemented yet"

	case inCalendar:
		return "Sorry, this is not implemented yet"

	case inAlert:
		return "Loco, implementa las alertas"
	}
	// por default se muestra el menu principal
	return docStyle.Render(m.mainMenu.View())
}

/*
triggered when and option is selected in the main menu.
Handles the App state and sets the correct mode
*/
func (a App) selectMode() (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	a.mainMenu.Selected = false
	// change app mode
	switch a.mainMenu.List.SelectedItem().FilterValue() {
	case "listaMats": // abrir la lista de materias entera
		a.Mode = inListMats
		var err error
		a.listaMats, err = NewListaMats(a.appHeight, a.appWith, a.config.FHorario)
        if err != nil {
            panic(err)
        }
		// if err != nil {
		// 	a.Mode = inAlert
		// }
		return a, cmd

	case "horario": // abrir mi horario TODO: IMPLEMENTAR
		a.Mode = inHorario
		return a, cmd

	case "calendario": // abrir el calendario TODO: IMPLEMENTAR
		a.Mode = inCalendar
		// a.Game = NewTyper(a.appWith)
		// cmd = a.Game.Init()
		return a, cmd

	case "nuevo_hor": // crear nuevo horario TODO: IMPLEMENTAR
		a.Mode = inSelection
		return a, cmd
	}
	return a, cmd
}
