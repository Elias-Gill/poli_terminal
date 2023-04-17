package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	armHors "github.com/elias-gill/poli_terminal/cli/armadorHorarios"
	"github.com/elias-gill/poli_terminal/cli/horario"
	cfman "github.com/elias-gill/poli_terminal/configManager"
	"github.com/elias-gill/poli_terminal/styles"
)

// modos
const (
	inMenu = iota
	inHorario
	inSelection
	inArmarHor
)

type App struct {
	Mode      int
	appWith   int
	appHeight int
	config    *cfman.Configurations

	// components
	mainMenu     MenuPrincipal
	horario      horario.DisplayHorario
	armador armHors.ArmadorHorario
}

func NewApp() App {
	return App{
		mainMenu: NewMainMenu(),
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

	case inArmarHor:
		a.armador, cmd = a.armador.Update(msg)
		if a.armador.Quit {
			a.mainMenu.List.SetWidth(a.appWith)
			a.mainMenu.List.SetHeight(a.appHeight)
			a.Mode = inMenu
		}

	case inMenu:
		a.mainMenu, cmd = a.mainMenu.Update(msg)
		if a.mainMenu.Selected {
			return a.selectMode()
		}
	}

	return a, cmd
}

// selecciona la vista dependiendo del estado de la aplicacion
func (m App) View() string {
	switch m.Mode {
	case inArmarHor:
		return styles.DocStyle.Render(m.armador.View())

	case inHorario:
		return styles.DocStyle.Render(m.horario.View())
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
	case "modHorario": // abrir la lista de materias entera
		a.Mode = inArmarHor
		a.armador = armHors.NewArmador()
		// truco para mandar informacion de tamano
		a.armador, _ = a.armador.Update(
			tea.WindowSizeMsg{
				Width:  a.appWith,
				Height: a.appHeight,
			},
		)

	case "horario": // abrir mi horario actual
		a.Mode = inHorario
		// TODO: continuar
		var err error
		a.horario = horario.NewHorario()
		if err != nil {
			panic(err)
		}

	case "calendario": // abrir el calendario
		// TODO: IMPLEMENTAR

	case "salir":
		return a, tea.Quit
	}
	return a, cmd
}
