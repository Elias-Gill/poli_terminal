package cli

import (
	tea "github.com/charmbracelet/bubbletea"
)

// modos
const (
	inMenu = iota
	inCalendar
	inHorario
	inSelection
	inListMats
)

type App struct {
	Mode      int
	appWith   int
	appHeight int
    fHorario  string
	// components
	mainMenu  MainMenu
	listaMats ListaMats
}

func NewApp() App {
	return App{
		mainMenu: NewMainMenu(),
        // TODO: leer de un archivo de configuracion
		fHorario: "/home/elias/Documentos/go_proyects/poli_terminal/assets/Horario-de-clases-y-examenes-del-Segundo-Periodo-2022-version-web-24032023.xlsx",
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
	case inCalendar: // el juego se encuentra corriendo
		// TODO: REALIZAR LOS CAMBIOS PERTINENTES
		// if a.Game.Done {
		// 	a.Game.Done = false
		// 	a.Mode = menu
		// 	return a, nil
		// }
		// actualizar el juego
		// a.Game, cmd = a.Game.Update(msg)
		return a, cmd

	case inListMats: // el juego se encuentra corriendo
		a.listaMats, cmd = a.listaMats.Update(msg)
		return a, cmd
	}

	a.mainMenu, cmd = a.mainMenu.Update(msg)
	// cambiar el modo despues de una seleccion
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
		a.listaMats = NewListaMats(a.appHeight, a.appWith, a.fHorario)
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
