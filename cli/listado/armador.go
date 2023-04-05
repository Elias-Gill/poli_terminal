package listado

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/excelParser"
	boxer "github.com/treilik/bubbleboxer"
)

const (
	selectAddr = "selector"
	infoAddr   = "informacion"
	listaAddr  = "listado"
)

func NewArmador(f string) Armador {
	// leaf content creation (models)
	info := newInfoMateria(excelParser.Materia{})
	selector, err := NewSelectorMats(f)
	if err != nil {
		panic("No se puede crear el selector de materias")
	}
	// lista := newInfoMateria(NewSelectorMats())

	// layout-tree defintion
	m := Armador{tui: boxer.Boxer{}}
	m.tui.LayoutTree = boxer.Node{
		// orientation
		VerticalStacked: true,
		// Los largos de los hijos, debe coincidir con la cantidad de nodos
		SizeFunc: func(_ boxer.Node, widthOrHeight int) []int {
			return []int{(widthOrHeight / 2) + 1, (widthOrHeight / 2) - 2}
		},
		Children: []boxer.Node{
			m.tui.CreateLeaf(selectAddr, selector),
			{
				SizeFunc: func(_ boxer.Node, widthOrHeight int) []int {
					return []int{widthOrHeight / 2, widthOrHeight / 2}
				},
				Children: []boxer.Node{
					m.tui.CreateLeaf(infoAddr, info),
					// m.tui.CreateLeaf("third", upper),
				},
			},
		},
	}

	return m
}

type Armador struct {
	tui  boxer.Boxer
	Quit bool
}

func (m Armador) Init() tea.Cmd {
	return spinner.Tick
}
func (m Armador) Update(msg tea.Msg) (Armador, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.tui.UpdateSize(msg)

	}
	return m, nil
}
func (m Armador) View() string {
	return m.tui.View()
}
