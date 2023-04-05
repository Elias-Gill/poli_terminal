package listado

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/excelParser"
	"github.com/elias-gill/poli_terminal/styles"
)

type infoMateria struct {
	Quit    bool
	materia excelParser.Materia
}

func newInfoMateria(m excelParser.Materia) infoMateria {
	return infoMateria{materia: m, Quit: false}
}

func (i infoMateria) Init() tea.Cmd { return nil }

func (i infoMateria) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	options := map[string]struct{}{"i": {}, "q": {}, "esc": {}}
	switch msg := msg.(type) {
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
	res := styles.DoneStyle.Render(i.materia.Nombre) + "\n"
	for i := 0; i < 50; i++ {
		res += "─"
	}
	res += "\n" + i.materia.Profesor + "\t" +
		"\n\n\t" + styles.GoodStyle.Render("Parcial 1: \t") + i.materia.Parcial1 + "\t" +
		"\n\t" + styles.GoodStyle.Render("Parcial 2: \t") + i.materia.Parcial2 + "\t" +
		"\n\n\t" + styles.GoodStyle.Render("Final 1: \t") + i.materia.Final1 + "\t" +
		"\n\t" + styles.GoodStyle.Render("Final 2: \t") + i.materia.Final2 + "\t"

	return styles.DocStyle.Render(res)
}
