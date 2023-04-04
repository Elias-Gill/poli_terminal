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

func (i infoMateria) Update(msg tea.Msg) (infoMateria, tea.Cmd) {
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
	var res string
	res += i.materia.Nombre + "\n" + i.materia.Profesor + "\n\n" +
		styles.GoodStyle.Render("Parcial 1: \t") + i.materia.Parcial1 +
		"\n" + styles.GoodStyle.Render("Parcial 2: \t") + i.materia.Parcial2 +
		"\n\n" + styles.GoodStyle.Render("Final 1: \t") + i.materia.Final1 +
		"\n" + styles.GoodStyle.Render("Final 2: \t") + i.materia.Final2

	return styles.DocStyle.Render(res)
}
