package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/cli/horario"
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
	inModHorario
)

type App struct {
	Mode      int
	appWith   int
	appHeight int
	config    configManager.Configurations

	// components
	mainMenu     MainMenu
	horario      horario.Horario
	selectorMats listado.Armador
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

	// handle events
	switch a.Mode {
	case inCalendar:
		// TODO: implementar

	case inHorario:
		a.horario, cmd = a.horario.Update(msg)
		if a.horario.Quit {
			a.Mode = inMenu
		}

	case inModHorario:
		a.selectorMats, cmd = a.selectorMats.Update(msg)
		if a.selectorMats.Quit {
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
	case inModHorario:
		return styles.DocStyle.Render(m.selectorMats.View())

	case inHorario:
		return styles.DocStyle.Render(m.horario.View())

	case inCalendar:
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
	case "modHorario": // abrir la lista de materias entera
		a.Mode = inModHorario
		a.selectorMats = listado.NewArmador(a.config.FHorario)
		// truco para mandar informacion de tamano
		a.selectorMats, _ = a.selectorMats.Update(
			tea.WindowSizeMsg{
				Width:  a.appWith,
				Height: a.appHeight,
			},
		)

    case "horario": // abrir mi horario actual TODO: continuar
		a.Mode = inHorario
		var err error
		a.horario = horario.NewHorario([]excelParser.Materia{
			{Nombre: "materoas1"},
			{Nombre: "materoas2"},
			{Nombre: "materoas3"},
		})
		if err != nil {
			panic(err)
		}

	case "calendario": // abrir el calendario TODO: IMPLEMENTAR
		// a.Mode = inCalendar

	}
	return a, cmd
}
