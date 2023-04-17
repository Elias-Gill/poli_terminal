package armadorHorarios

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	pts "github.com/elias-gill/poli_terminal/cli/prompts"
	"github.com/elias-gill/poli_terminal/configManager"
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
	inPrompt
)

type ArmadorHorario struct {
	width       int
	height      int
	Quit        bool
	mode        int
	prompt      pts.Prompt
	infoMat     *infoMateria
	listaSelecs *listSelecs
	selector    *selectMats
}

func NewArmador() ArmadorHorario {
	c := configManager.GetUserConfig()
	selector := newSelectorMats()
	return ArmadorHorario{
		mode:        inSelector,
		Quit:        false,
		infoMat:     newInfoMateria(selector.Focused),
		listaSelecs: newListaSelecciones(c.MateriasUsuario),
		selector:    selector,
	}
}

func (a ArmadorHorario) Init() tea.Cmd { return nil }

func (a ArmadorHorario) Update(msg tea.Msg) (ArmadorHorario, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// salir
		if msg.String() == "q" || msg.String() == tea.KeyEsc.String() {
			// TODO:delegar al selector
			if !a.selector.Filtering || a.mode == inLista {
				a.mode = inPrompt
				a.prompt = pts.NewPrompt("Desea GUARDAR este nuevo horario ?")
				return a, nil
			}
		}

		// mensaje de confirmacion
		// TODO: refactor
		if a.mode == inPrompt {
			if msg.String() == "enter" {
				if a.prompt.Selection == "Yes" {
					c := configManager.GetUserConfig()
					c.ChangeMateriasUsuario(a.listaSelecs.lista)
					c.WriteUserConfig()
				}
				a.Quit = true
				return a, nil
			}
			a.prompt = a.prompt.Update(msg)
			return a, nil
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
		if a.mode == inPrompt {
			a.prompt = a.prompt.Update(msg)
			return a, nil
		}
		return a.UpdateSize(msg), nil
	}

	if a.mode == inSelector {
		// anadir materia
		cmd = a.selector.Update(msg)
		a.infoMat.ChangeMateria(a.selector.Focused)
		if !a.selector.Filtering {
			if a.selector.IsSelected {
				a.listaSelecs.AddMateria(a.selector.Focused)
				a.selector.IsSelected = false
			}
		}
		return a, cmd
	}

	if a.mode == inLista {
		cmd = a.listaSelecs.Update(msg)
		return a, cmd
	}

	return a, cmd
}

func (a ArmadorHorario) View() string {
	if a.mode == inPrompt {
		return a.prompt.View()
	}
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

	a.infoMat.Update(info)
	a.selector.Update(selector)
	a.listaSelecs.Update(lista)

	return a
}
