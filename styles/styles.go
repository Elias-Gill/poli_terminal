package styles

import "github.com/charmbracelet/lipgloss"

var (
	// docStyle corresponde al estilo utilizado para renderizar la app entera
	DocStyle = lipgloss.NewStyle().Width(200).
			Height(20).
			Margin(1, 2)

	// styles to colorate strings while the user is typing
	GoodStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#43d11c"))

	BadStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#f40045"))

	DoneStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#ae67f0"))

	AuthorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#999999"))
)
