package cli

import tea "github.com/charmbracelet/bubbletea"

type App struct {
	Quit      bool
	Mode      int
	appWith   int
	appHeight int
}

func NewApp() App {
	return App{
		Quit: false,
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
		a.appWith = msg.Height
	}

	// handle events
	switch a.Mode {
	case 1: // el juego se encuentra corriendo
	}
	return a, cmd
}

// selecciona la vista dependiendo del estado de la aplicacion
func (m App) View() string {
	switch m.Mode {
	case 1: // mostrar juego
	}
	// return docStyle.Render(m.Menu.View())
	return ""
}

/*
	triggered when and option is selected in the main menu.

Handles the App state and sets the correct mode
*/
func (a App) selectMode() (tea.Model, tea.Cmd) {
	return a, nil
}

func (a App) selectLocalQuote() (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return a, cmd
}
