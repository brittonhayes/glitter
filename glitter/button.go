package glitter

import (
	"github.com/charmbracelet/lipgloss"
)

// ButtonStyle maps
// each button aesthetic to
// an integer
type ButtonStyle int

const (
	// Info is a button with a
	// primary foreground color
	Info ButtonStyle = iota

	// Warn is a button with a
	// yellow to indicate an issue
	Warn

	// Error is a button with a
	// red color to indicate a critical
	// problem
	Error

	// Success is a button with a
	// green color to indicate
	// everything is good
	Success
)

// Button is a button shaped UI element with a few different
// style options
func (ui *UI) Button(body string, style ButtonStyle) lipgloss.Style {
	switch style {
	case Info:
		return ui.button().Copy().Background(ui.Theme.Normal.Cyan).Foreground(ui.Theme.Primary.Background).SetString(body)
	case Warn:
		return ui.button().Copy().Background(ui.Theme.Bright.Yellow).Foreground(ui.Theme.Primary.Background).SetString(body)
	case Error:
		return ui.button().Copy().Background(ui.Theme.Bright.Red).Foreground(ui.Theme.Primary.Background).SetString(body)
	case Success:
		return ui.button().Copy().Background(ui.Theme.Bright.Green).Foreground(ui.Theme.Primary.Background).SetString(body)
	default:
		return ui.button().SetString(body)
	}
}

func (ui *UI) button() lipgloss.Style {
	return lipgloss.NewStyle().
		Padding(0, 1)
}
