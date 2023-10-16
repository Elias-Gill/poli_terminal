package schedule

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	ep "github.com/elias-gill/poli_terminal/excelParser"
	"github.com/elias-gill/poli_terminal/styles"
)

type subjectInfo struct {
	materia *ep.Materia
	width   int
	height  int
}

func newInfoMateria(m *ep.Materia) *subjectInfo {
	return &subjectInfo{materia: m}
}

func (i *subjectInfo) Init() tea.Cmd { return nil }

func (i *subjectInfo) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		i.width = msg.Width
		i.height = msg.Height
	}
	return nil
}

func (i *subjectInfo) View() string {
	if i.materia == nil {
		return "\n\n\n\n\n\n\n"
	}
	res := styles.DoneStyle.Render(i.materia.Profesor) +
		"\n" + styles.GoodStyle.Render("Parciales: \t") +
		"\n" + i.materia.Parcial1 +
		"\n" + i.materia.Parcial2 +
		"\n" + styles.GoodStyle.Render("Finales: \t") +
		"\n" + i.materia.Final1 +
		"\n" + i.materia.Final2 + "\n"

	return lipgloss.PlaceHorizontal(i.width, lipgloss.Left, res)
}

func (i *subjectInfo) CambiarMateria(m *ep.Materia) {
	i.materia = m
}
