package armadorHorarios

import (
	tea "github.com/charmbracelet/bubbletea"
	ep "github.com/elias-gill/poli_terminal/excelParser"
	"github.com/elias-gill/poli_terminal/styles"
	boxer "github.com/treilik/bubbleboxer"
)

const (
	selectAddr = "selector"
	infoAddr   = "informacion"
	listaAddr  = "listado"
)
const (
	inSelector = iota
	inLista
)

type ArmadorHorario struct {
	tui      boxer.Boxer
	Quit     bool
	mode     int
	materias []ep.Materia
}

func NewArmador(f string) ArmadorHorario {
	// modelos
	info := newInfoMateria(ep.Materia{})
	lista := NewLista([]ep.Materia{})
	selector, err := NewSelectorMats(f)
	if err != nil {
		panic("No se puede crear el selector de materias")
	}

	// layout-tree defintion
	m := ArmadorHorario{tui: boxer.Boxer{}, mode: inSelector}
	m.tui.LayoutTree = boxer.Node{
		// orientation
		// Los largos de los hijos, debe coincidir con la cantidad de nodos
		SizeFunc: func(_ boxer.Node, widthOrHeight int) []int {
			return []int{(widthOrHeight * 3 / 5), (widthOrHeight * 2 / 5)}
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
					m.tui.CreateLeaf(listaAddr, lista),
				},
			},
		},
	}

	return m
}

func (a ArmadorHorario) Init() tea.Cmd { return nil }

func (a ArmadorHorario) Update(msg tea.Msg) (ArmadorHorario, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
        case "q":
            a.Quit = true

        // cambiar de modo
		case tea.KeyTab.String():
			if a.mode == inLista {
				a.mode = inSelector
			} else {
				a.mode = inLista
			}
		}

	case tea.WindowSizeMsg:
		h, w := styles.DocStyle.GetFrameSize()
		msg.Height -= h
		msg.Width -= w
		a.tui.UpdateSize(msg)
	}

	// si estamos en la lista de seleccionados
	if a.mode == inLista {
		a.tui.ModelMap[listaAddr], cmd = a.tui.ModelMap[listaAddr].Update(msg)
		return a, cmd
	}

	// actualizar el selector
	a.tui.ModelMap[selectAddr], cmd = a.tui.ModelMap[selectAddr].Update(msg)
	elem := a.tui.ModelMap[selectAddr]
	// traer la materia seleccionada
	switch selec := elem.(type) {
	case SelectMats:
		// actualizar la info
		a.tui.ModelMap[infoAddr] = newInfoMateria(selec.Focus)
		// si se preciono enter, agregar materia
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "q", "esc":
				if !selec.Filtering {
					a.Quit = true
					return a, nil
				}

			case "enter":
				switch lmats := a.tui.ModelMap[listaAddr].(type) {
				case listSelecs:
					a.tui.ModelMap[listaAddr] = lmats.AddMateria(selec.Focus)
				}
			}
		}
	}
	return a, cmd
}

func (a ArmadorHorario) View() string {
	return a.tui.View()
}
