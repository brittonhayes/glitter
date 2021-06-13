package glitter

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

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

func (ui *UI) Comment(prefix string, body string) lipgloss.Style {
	comment := lipgloss.NewStyle().
		Foreground(ui.Theme.Dim.White).
		Faint(true).
		SetString(fmt.Sprintf("%s %s", prefix, body))

	return comment
}
