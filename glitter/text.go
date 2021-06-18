package glitter

import (
	"github.com/charmbracelet/lipgloss"
)

// Italicize puts a string into italics and applies
// the UI theme
func (ui *UI) Italicize(body string) lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Italic(true).
		Foreground(ui.Theme.Primary.Foreground).
		SetString(body)
}

// Accent highlights a text element as important
// by setting it bold and bright cyan
func (ui *UI) Accent(body string) lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(ui.Theme.Bright.Cyan).
		SetString(body)
}
