package configManager

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/knipferrc/teacup/filetree"
)

// Filetree represents the properties of the UI.
type Filetree struct {
	filetree filetree.Bubble
}

// NewFiletree creates a new instance of the UI.
func NewFiletree() Filetree {
	filetreeModel := filetree.New(
		true,
		true,
		"",
		"",
		lipgloss.AdaptiveColor{Light: "#000000", Dark: "63"},
		lipgloss.AdaptiveColor{Light: "#000000", Dark: "63"},
		lipgloss.AdaptiveColor{Light: "63", Dark: "63"},
		lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
	)

	return Filetree{
		filetree: filetreeModel,
	}
}

// Init intializes the UI.
func (b Filetree) Init() tea.Cmd {
	return b.filetree.Init()
}

// Update handles all UI interactions.
func (b Filetree) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.filetree.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			cmds = append(cmds, tea.Quit)
		}
	}

	b.filetree, cmd = b.filetree.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}

// View returns a string representation of the UI.
func (b Filetree) View() string {
	return b.filetree.View()
}
