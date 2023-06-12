package prompts

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ConfirmPrompt struct {
	Selection bool
	Quit      bool
	Msg       string
}

func NewConfirmPrompt(msg string) *ConfirmPrompt {
	return &ConfirmPrompt{
		Msg:       msg,
		Selection: false,
		Quit:      false,
	}
}

func (p *ConfirmPrompt) Init() tea.Cmd { return nil }

func (p *ConfirmPrompt) Update(msg tea.Msg) *ConfirmPrompt {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			p.Quit = true
		}

		if msg.String() == "q" {
			p.Selection = false
			p.Quit = true
		}

		options := map[string]struct{}{"j": {}, "k": {}, "h": {}, "l": {}, "left": {}, "right": {}}
		if _, ok := options[msg.String()]; ok {
			p.Selection = !p.Selection
		}
	}
	return p
}

func (p *ConfirmPrompt) View() string {
	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Border(lipgloss.RoundedBorder(), true)

	tittle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Border(lipgloss.RoundedBorder(), false, false, true)

	selection := lipgloss.NewStyle().Foreground(lipgloss.Color("#ae67f0"))

	yes := "Yes"
	no := "No"
	if p.Selection {
		yes = selection.Render(yes)
	} else {
		no = selection.Render(no)
	}
	return style.Render(
		tittle.Render(p.Msg+"\n") + "\n" + yes + "    " + no)
}
