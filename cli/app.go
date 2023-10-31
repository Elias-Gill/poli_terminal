package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	consts "github.com/elias-gill/poli_terminal/cli/constants"
	"github.com/elias-gill/poli_terminal/cli/menus"
	"github.com/elias-gill/poli_terminal/cli/schedule"
	"github.com/elias-gill/poli_terminal/cli/schedule/armador_horario"
	cfman "github.com/elias-gill/poli_terminal/configManager"
	"github.com/elias-gill/poli_terminal/styles"
)

type App struct {
	Mode       int
	appWith    int
	appHeight  int
	config     *cfman.Configurations
	components map[int]consts.Component
}

func NewApp() App {
	return App{
		config: cfman.GetUserConfig(),
		Mode:   consts.InMainMenu,
		// components
		components: map[int]consts.Component{
			consts.InMainMenu:          menus.NewMainMenu(),
			consts.InScheduleDisplayer: schedule.NewHorario(),
			consts.InConfigMenu:        menus.NewConfigMenu(),
			consts.InScheduleMaker:     armador_horario.NewArmadorHorario(),
		},
	}
}

func (a App) Init() tea.Cmd {
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

	a.components[a.Mode], cmd = a.components[a.Mode].Update(msg)
    a.Mode = consts.CurrentMode
	return a, cmd
}

// selecciona la vista dependiendo del estado de la aplicacion
func (a App) View() string {
	return styles.DocStyle.Render(a.components[a.Mode].Render())
}
