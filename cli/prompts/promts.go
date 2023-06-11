package prompts

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Prompt struct {
	Selection  bool
	Quit      bool
	Msg       string
}

func NewPrompt(msg string) *Prompt {
	return &Prompt{
		Msg:       msg,
		Selection:  false,
		Quit:      false,
	}
}

func (p *Prompt) Init() tea.Cmd { return nil }

func (p *Prompt) Update(msg tea.Msg) *Prompt {
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

// TODO: refactor
func (p *Prompt) View() string {
	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Border(lipgloss.RoundedBorder(), true)

	tittle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Border(lipgloss.RoundedBorder(), true)

	selection := lipgloss.NewStyle().Foreground(lipgloss.Color("#912"))

	yes := "Yes"
	no := "No"
	if p.Selection {
		yes = selection.Render(yes)
	} else {
		no = selection.Render(no)
	}
	return style.Render(
		tittle.Render(p.Msg+"\n") + "\n" + yes + "\t" + no)
}
