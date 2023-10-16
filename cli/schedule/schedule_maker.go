package schedule

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elias-gill/poli_terminal/cli/constants"
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

type ScheduleMaker struct {
	width           int
	height          int
	mode            int
	prompt          *pts.ConfirmPrompt
	infoMat         *subjectInfo
	selectedList    *selectedList
	subjectSelector *selector
}

func NewScheduleMaker() ScheduleMaker {
	selector, err := newSelectorMats()
	if err != nil {
		panic("No se pudo inicializar el armador de horarios")
	}
	inf := newInfoMateria(selector.Focused)
	lista := newListaSelecciones()
	return ScheduleMaker{
		mode:            inSelector,
		infoMat:         inf,
		selectedList:    lista,
		subjectSelector: selector,
	}
}

func (a ScheduleMaker) Init() tea.Cmd { return nil }

func (a ScheduleMaker) Update(msg tea.Msg) (constants.Component, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == tea.KeyEsc.String() {
			if a.mode != inSelector || !a.subjectSelector.Filtering {
				a.mode = inPrompt
				a.prompt = pts.NewConfirmPrompt("Desea GUARDAR este nuevo horario ?")
				return a, nil
			}
		}

		// prompt de confirmacion
		if a.mode == inPrompt {
			a.prompt.Update(msg)
			if a.prompt.Quit {
				if a.prompt.Selection {
					c := configManager.GetUserConfig()
					c.ChangeMateriasUsuario(a.selectedList.lista)
					c.WriteUserConfig()
				}
                constants.CurrentMode = constants.InMainMenu
			}
			return a, nil
		}

		// cambiar de modo (tab)
		if msg.String() == tea.KeyTab.String() {
			if a.mode == inLista {
				a.mode = inSelector
			} else {
				a.mode = inLista
			}
			a.selectedList.isFocused = !a.selectedList.isFocused
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
		cmd = a.subjectSelector.Update(msg)
		a.infoMat.CambiarMateria(a.subjectSelector.Focused)
		if a.subjectSelector.IsSelected {
			a.selectedList.AddMateria(a.subjectSelector.Focused)
			a.subjectSelector.IsSelected = false
		}
		return a, cmd
	}

	if a.mode == inLista {
		cmd = a.selectedList.Update(msg)
		return a, cmd
	}

	return a, cmd
}

func (a ScheduleMaker) Render() string {
	if a.mode == inPrompt {
		return a.prompt.View()
	}
	aux := lipgloss.JoinVertical(0, a.infoMat.View(), a.selectedList.View())
	return lipgloss.JoinHorizontal(0, a.subjectSelector.View(), aux)
}

// calcula los tamanos necesarios para los objetos en pantalla
func (a ScheduleMaker) UpdateSize(m tea.WindowSizeMsg) ScheduleMaker {
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
	a.subjectSelector.Update(selector)
	a.selectedList.Update(lista)

	return a
}
