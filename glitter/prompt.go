package glitter

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)


// Prompt is a terminal font that allows for a custom prefix
// and command body
//
// This works well for CLI tutorials and example commands
//
func (ui *UI) Prompt(prefix string, body string, faint bool) (lipgloss.Style, lipgloss.Style) {
	var (
		p = lipgloss.NewStyle().
			Bold(true).
			Foreground(ui.Theme.Bright.Green).
			Faint(faint).
			SetString(prefix)
		b = lipgloss.NewStyle().
			Foreground(ui.Theme.Bright.White).
			Faint(faint).
			SetString(body)
	)

	return p, b
}

// Comment is code comment that renders a lighter tone
// comment string to the screen
//
// This works well for using before your prompt elements
func (ui *UI) Comment(prefix string, body string) lipgloss.Style {
	comment := lipgloss.NewStyle().
		Foreground(ui.Theme.Dim.White).
		Faint(true).
		SetString(fmt.Sprintf("%s %s", prefix, body))

	return comment
}
