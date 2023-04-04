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
	inListMats
)

type App struct {
	Mode      int
	appWith   int
	appHeight int
	config    configManager.Configurations
	// components
	mainMenu     MainMenu
	listaMats    horario.Horario
	selectorMats *listado.SelectMats
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
		a.listaMats, cmd = a.listaMats.Update(msg)
		if a.listaMats.Quit {
			a.Mode = inMenu
		}

	case inListMats:
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
	case inListMats:
		return styles.DocStyle.Render(m.selectorMats.View())

	case inHorario:
		return styles.DocStyle.Render(m.listaMats.View())

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
		a.Mode = inListMats
		var err error
		a.selectorMats, err = listado.NewSelectorMats(a.appHeight, a.appWith, a.config.FHorario)
		if err != nil {
			panic(err)
		}

	case "horario": // abrir mi horario actual
		a.Mode = inHorario
		var err error
		a.listaMats = horario.NewHorario([]excelParser.Materia{
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
