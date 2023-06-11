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
	inMainMenu = iota
	inScheduleDisplayer
	inConfigMenu
	inCalendar
	inScheduleMaker
)

type App struct {
	Mode      int
	appWith   int
	appHeight int
	config    *cfman.Configurations

	// components
	mainMenu   menus.MainMenu
	configMenu menus.ConfigMenu
	displayer  schedule.ScheduleDisplayer
	maker      schedule.ScheduleMaker
}

func NewApp() App {
	return App{
		mainMenu: menus.NewMainMenu(),
		config:   cfman.GetUserConfig(),
		Mode:     inMainMenu,
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
	case inScheduleDisplayer:
		a.displayer, cmd = a.displayer.Update(msg)
		if a.displayer.Quit {
			a.mainMenu.List.SetWidth(a.appWith)
			a.mainMenu.List.SetHeight(a.appHeight)
			a.Mode = inMainMenu
		}

	case inScheduleMaker:
		a.maker, cmd = a.maker.Update(msg)
		if a.maker.Quit {
			a.mainMenu.List.SetWidth(a.appWith)
			a.mainMenu.List.SetHeight(a.appHeight)
			a.Mode = inMainMenu
		}

	case inMainMenu:
		a.mainMenu, cmd = a.mainMenu.Update(msg)
		if a.mainMenu.IsSelected {
			return a.selectMode()
		}

	case inConfigMenu:
		a.configMenu, cmd = a.configMenu.Update(msg)
		if a.configMenu.Quit {
			a.mainMenu.List.SetWidth(a.appWith)
			a.mainMenu.List.SetHeight(a.appHeight)
			a.Mode = inMainMenu
		}
	}

	return a, cmd
}

// selecciona la vista dependiendo del estado de la aplicacion
func (m App) View() string {
	switch m.Mode {
	case inScheduleMaker:
		return styles.DocStyle.Render(m.maker.View())

	case inScheduleDisplayer:
		return styles.DocStyle.Render(m.displayer.View())

	case inConfigMenu:
		return styles.DocStyle.Render(m.configMenu.View())
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
		a.maker, _ = a.maker.Update(
			tea.WindowSizeMsg{
				Width:  a.appWith,
				Height: a.appHeight,
			},
		)

	case "horario": // abrir la vista del horario
		a.Mode = inScheduleDisplayer
		a.displayer = schedule.NewScheduleDisplayer()
		a.displayer, _ = a.displayer.Update(
			tea.WindowSizeMsg{
				Width:  a.appWith,
				Height: a.appHeight,
			},
		)

	case "calendario": // abrir el calendario
		// TODO: IMPLEMENTAR

	case "configMenu": // abrir el menu de configuracion
		a.Mode = inConfigMenu
		a.configMenu = menus.NewConfigMenu()
		a.configMenu, _ = a.configMenu.Update(
			tea.WindowSizeMsg{
				Width:  a.appWith,
				Height: a.appHeight,
			},
		)

	case "salir":
		return a, tea.Quit
	}
	return a, cmd
}
