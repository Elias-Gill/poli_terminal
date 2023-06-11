package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/cli/menus"
	"github.com/elias-gill/poli_terminal/cli/schedule"
	cfman "github.com/elias-gill/poli_terminal/configManager"
	"github.com/elias-gill/poli_terminal/styles"
)

// modos
const (
	inMenu = iota
	inHorario
	inSelection
	inScheduleMaker
)

type App struct {
	Mode      int
	appWith   int
	appHeight int
	config    *cfman.Configurations

	// components
	mainMenu menus.MainMenu
	horario  schedule.ScheduleDisplayer
	maker    schedule.ScheduleMaker
}

func NewApp() App {
	return App{
		mainMenu: menus.NewMainMenu(),
		config:   cfman.GetUserConfig(),
		Mode:     inMenu,
	}
}

func (m App) Init() tea.Cmd {
	return nil
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	// actualizar el tamano de la app
	case tea.WindowSizeMsg:
		a.appWith = msg.Width
		a.appHeight = msg.Height

	// salir
	case tea.KeyMsg:
		if msg.String() == tea.KeyCtrlC.String() {
			return a, tea.Quit
		}
	}

	// handle events per mode
	switch a.Mode {
	case inHorario:
		a.horario, cmd = a.horario.Update(msg)
		if a.horario.Quit {
			a.mainMenu.List.SetWidth(a.appWith)
			a.mainMenu.List.SetHeight(a.appHeight)
			a.Mode = inMenu
		}

	case inScheduleMaker:
		a.maker, cmd = a.maker.Update(msg)
		if a.maker.Quit {
			a.mainMenu.List.SetWidth(a.appWith)
			a.mainMenu.List.SetHeight(a.appHeight)
			a.Mode = inMenu
		}

	case inMenu:
		a.mainMenu, cmd = a.mainMenu.Update(msg)
		if a.mainMenu.IsSelected {
			return a.selectMode()
		}
	}

	return a, cmd
}

// selecciona la vista dependiendo del estado de la aplicacion
func (m App) View() string {
	switch m.Mode {
	case inScheduleMaker:
		return styles.DocStyle.Render(m.maker.View())

	case inHorario:
		return styles.DocStyle.Render(m.horario.View())
	}
	// por default se muestra el menu principal
	return styles.DocStyle.Render(m.mainMenu.View())
}

/*
triggered when an option is selected in the main menu.
Handles the App state and sets the correct mode
*/
func (a App) selectMode() (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	a.mainMenu.IsSelected = false
	// change app mode
	switch a.mainMenu.List.SelectedItem().FilterValue() {
	case "scheduleMaker":
		a.Mode = inScheduleMaker
		a.maker = schedule.NewScheduleMaker()
		// truco para mandar informacion de tamano
		a.maker, _ = a.maker.Update(
			tea.WindowSizeMsg{
				Width:  a.appWith,
				Height: a.appHeight,
			},
		)

	case "horario": // abrir mi horario actual TODO: continuar
		a.Mode = inHorario
		var err error
		a.horario = schedule.NewScheduleDisplayer()
		if err != nil {
			panic(err)
		}

	case "calendario": // abrir el calendario
		// TODO: IMPLEMENTAR

	case "configMenu": // abrir el calendario
		// TODO: IMPLEMENTAR

	case "salir":
		return a, tea.Quit
	}
	return a, cmd
}
