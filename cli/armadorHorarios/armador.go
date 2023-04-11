package armadorHorarios

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	ep "github.com/elias-gill/poli_terminal/excelParser"
	"github.com/elias-gill/poli_terminal/styles"
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
	width       int
	height      int
	Quit        bool
	mode        int
	infoMat     infoMateria
	listaSelecs listSelecs
	selector    SelectMats
}

func NewArmador(f string) ArmadorHorario {
	return ArmadorHorario{
		mode:        inSelector,
		Quit:        false,
		infoMat:     newInfoMateria(ep.Materia{}),
		listaSelecs: newLista([]ep.Materia{}),
		selector:    newSelectorMats(f),
	}
}

func (a ArmadorHorario) Init() tea.Cmd { return nil }

func (a ArmadorHorario) Update(msg tea.Msg) (ArmadorHorario, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// salir
		if msg.String() == "q" {
			if !a.selector.Filtering || a.mode == inLista {
				a.Quit = true
				return a, nil
			}
		}

		// cambiar de modo (tab)
		if msg.String() == tea.KeyTab.String() {
			if a.mode == inLista {
				a.mode = inSelector
                a.listaSelecs.color = 2
			} else {
				a.mode = inLista
                a.listaSelecs.color = 1
			}
			return a, nil
		}

	case tea.WindowSizeMsg:
		return a.UpdateSize(msg), nil
	}

	if a.mode == inSelector {
		// anadir materia con enter
		a.selector, cmd = a.selector.Update(msg)
		if !a.selector.Filtering {
			if a.selector.Selected {
				a.listaSelecs = a.listaSelecs.AddMateria(a.selector.Focused)
				a.selector.Selected = false
			}
			a.infoMat = a.infoMat.ChangeMateria(a.selector.Focused)
		}
		return a, cmd
	}

	if a.mode == inLista {
		a.listaSelecs, cmd = a.listaSelecs.Update(msg)
		return a, cmd
	}

	return a, cmd
}

func (a ArmadorHorario) View() string {
	aux := lipgloss.JoinVertical(0, a.infoMat.View(), a.listaSelecs.View())
	return lipgloss.JoinHorizontal(0, a.selector.View(), aux)
}

// calcula los tamanos necesarios para los objetos en pantalla
func (a ArmadorHorario) UpdateSize(m tea.WindowSizeMsg) ArmadorHorario {
	var selector, info, lista tea.WindowSizeMsg
	x, y := styles.DocStyle.GetFrameSize()
	m.Height -= y
	m.Width -= x

	selector.Width = m.Width / 2
	selector.Height = m.Height

	info.Width = m.Width / 2
	info.Height = m.Height / 2

	lista.Width = m.Width / 2
	lista.Height = m.Height / 2

	a.infoMat, _ = a.infoMat.Update(info)
	a.selector, _ = a.selector.Update(selector)
	a.listaSelecs, _ = a.listaSelecs.Update(lista)

	return a
}
