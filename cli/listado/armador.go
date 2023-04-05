package listado

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elias-gill/poli_terminal/excelParser"
	"github.com/elias-gill/poli_terminal/styles"
	boxer "github.com/treilik/bubbleboxer"
)

const (
	selectAddr = "selector"
	infoAddr   = "informacion"
	listaAddr  = "listado"
)

type Armador struct {
	tui           boxer.Boxer
	Quit          bool
	selectorFocus bool
}

func NewArmador(f string) Armador {
	// leaf content creation (models)
	info := newInfoMateria(excelParser.Materia{})
	selector, err := NewSelectorMats(f)
	if err != nil {
		panic("No se puede crear el selector de materias")
	}
	// lista := newInfoMateria(NewSelectorMats())
	m := Armador{
		tui: boxer.Boxer{}, selectorFocus: true,
	}

	// layout-tree defintion
	m.tui.LayoutTree = boxer.Node{
		// orientation
		// Los largos de los hijos, debe coincidir con la cantidad de nodos
		SizeFunc: func(_ boxer.Node, widthOrHeight int) []int {
			return []int{(widthOrHeight / 2), (widthOrHeight / 2)}
		},
		Children: []boxer.Node{
			// hijo 1
			m.tui.CreateLeaf(selectAddr, selector),
			// hijo 2
			{
				VerticalStacked: true,
				SizeFunc: func(_ boxer.Node, widthOrHeight int) []int {
					return []int{(widthOrHeight / 2) - 1, (widthOrHeight / 2) + 1}
				},
				Children: []boxer.Node{
					m.tui.CreateLeaf(infoAddr, info),
					m.tui.CreateLeaf(listaAddr, info), // TODO: materias seleccionadas
				},
			},
		},
	}

	return m
}

func (m Armador) Init() tea.Cmd {
	return spinner.Tick
}
func (m Armador) Update(msg tea.Msg) (Armador, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			m.Quit = true
			return m, nil
		case tea.KeyTab.String():
			m.selectorFocus = !m.selectorFocus
		}

	case tea.WindowSizeMsg:
		h, w := styles.DocStyle.GetFrameSize()
		msg.Height -= h
		msg.Width -= w
		m.tui.UpdateSize(msg)
	}

	// si estamos en la lista de seleccionados
	if !m.selectorFocus {
		m.tui.ModelMap[listaAddr], cmd = m.tui.ModelMap[listaAddr].Update(msg)
		return m, cmd
	}

	// si estamos en el selector de materias
	m.tui.ModelMap[selectAddr], cmd = m.tui.ModelMap[selectAddr].Update(msg)
	aux := m.tui.ModelMap[selectAddr]
	// truquito para traer la materia seleccionada
	switch s := aux.(type) {
	case SelectMats:
		m.tui.ModelMap[infoAddr] = newInfoMateria(s.focused)
	}

	return m, cmd
}
func (m Armador) View() string {
	return m.tui.View()
}
