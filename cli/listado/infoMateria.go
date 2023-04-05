package listado

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/excelParser"
	"github.com/elias-gill/poli_terminal/styles"
)

type infoMateria struct {
	Quit    bool
	materia excelParser.Materia
	width   int
	height  int
}

func newInfoMateria(m excelParser.Materia) infoMateria {
	return infoMateria{materia: m, Quit: false}
}

func (i infoMateria) Init() tea.Cmd { return nil }

func (i infoMateria) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	options := map[string]struct{}{"i": {}, "q": {}, "esc": {}}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		i.width = msg.Width
		i.height = msg.Height

	case tea.KeyMsg:
		// si la tecla precionada es una de las de salir
		_, keyExit := options[msg.String()]
		if keyExit {
			i.Quit = true
		}
	}
	return i, nil
}

func (i infoMateria) View() string {
	res := styles.DoneStyle.Render(i.materia.Profesor) +
		"\n" + styles.GoodStyle.Render("Parciales: \t") +
		"\n" + i.materia.Parcial1 +
		"\n" + i.materia.Parcial2 +
		"\n" + styles.GoodStyle.Render("Finales: \t") +
		"\n" + i.materia.Final1 +
		"\n" + i.materia.Final2

	return res
}
