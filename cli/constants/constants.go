package constants

import tea "github.com/charmbracelet/bubbletea"

// modos
const (
	InMainMenu = iota
	InScheduleDisplayer
	InConfigMenu
	InCalendar
	InScheduleMaker
)

var CurrentMode int = InMainMenu

type Component interface {
	Update(msg tea.Msg) (Component, tea.Cmd)
	Render() string
}
