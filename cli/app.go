package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/cli/listado"
	"github.com/elias-gill/poli_terminal/configManager"
	"github.com/elias-gill/poli_terminal/excelParser"
	"github.com/elias-gill/poli_terminal/styles"
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
	horario   listado.Horario
	listaMats *listado.ListaMats
}

func NewApp() App {
	return App{
		mainMenu: NewMainMenu(),
		config:   configManager.GetUserConfig(),
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
	case inCalendar:
		// TODO: implementar

	case inAlert:
		// TODO: implementar

	case inHorario:
        a.horario, cmd = a.horario.Update(msg)
		if a.horario.Quit {
			a.Mode = inMenu
		}

	case inListMats:
		a.listaMats, cmd = a.listaMats.Update(msg)
		if a.listaMats.Quit {
			a.Mode = inMenu
		}

	case inMenu:
		// por defecto nos encontramos en el menu principal
		a.mainMenu, cmd = a.mainMenu.Update(msg)
		// WARN: no tratar de refactorear, problemas de performance
		if a.mainMenu.Selected {
			return a.selectMode()
		}
	}

	return a, cmd
}

// selecciona la vista dependiendo del estado de la aplicacion
func (m App) View() string {
	switch m.Mode {
	case inListMats:
		return styles.DocStyle.Render(m.listaMats.View())

	case inHorario:
		return styles.DocStyle.Render(m.horario.View())

	case inCalendar:
		// TODO: IMPLEMENTAR

	case inAlert:
		// TODO: IMPLEMENTAR
	}
	// por default se muestra el menu principal
	return styles.DocStyle.Render(m.mainMenu.View())
}

/*
triggered when and option is 3elected in the main menu.
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
		a.listaMats, err = listado.NewListaMats(a.appHeight, a.appWith, a.config.FHorario)
		if err != nil {
			panic(err)
		}

	case "horario": // abrir mi horario
		a.Mode = inHorario
		var err error
		a.horario = listado.NewInfoMateria([]excelParser.Materia{
			{Nombre: "materoas1"},
			{Nombre: "materoas2"},
			{Nombre: "materoas3"},
		})
		if err != nil {
			panic(err)
		}

	case "calendario": // abrir el calendario TODO: IMPLEMENTAR
		// a.Mode = inCalendar

	case "nuevo_hor": // crear nuevo horario TODO: IMPLEMENTAR
		// a.Mode = inSelection
	}
	return a, cmd
}
