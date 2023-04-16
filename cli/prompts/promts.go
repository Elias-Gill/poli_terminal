package prompts

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Prompt struct {
	Selection string
	Selected  bool
	Quit      bool
	Msg       string
}

func NewPrompt(msg string) Prompt {
	return Prompt{
		Selection: "Yes",
		Msg:       msg,
		Selected:  false,
		Quit:      false,
	}
}

func (p Prompt) Init() tea.Cmd { return nil }

func (p Prompt) Update(msg tea.Msg) Prompt {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			p.Quit = true
			p.Selected = true
		}

		if msg.String() == "q" {
			p.Quit = true
		}

		options := map[string]struct{}{"j": {}, "k": {}, "h": {}, "l": {}, "left": {}, "right": {}}
		if _, ok := options[msg.String()]; ok {
			if p.Selection == "Yes" {
				p.Selection = "No"
			} else {
				p.Selection = "Yes"
			}
		}
	}
	return p
}

func (p Prompt) View() string {
	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Border(lipgloss.RoundedBorder(), true)

	tittle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Border(lipgloss.RoundedBorder(), true)

	selection := lipgloss.NewStyle().Foreground(lipgloss.Color("#912"))

	yes := "Yes"
	no := "No"
	if p.Selection == "Yes" {
		yes = selection.Render(yes)
	} else {
		no = selection.Render(no)
	}
	return style.Render(
		tittle.Render(p.Msg+"\n") + "\n" + yes + "\t" + no)
}
