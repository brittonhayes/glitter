package glitter

import (
	"github.com/charmbracelet/lipgloss"
)

type ButtonStyle int

const (
	Info ButtonStyle = iota
	Warn
	Error
	Success
)

func (ui *UI) Button(body string, style ButtonStyle) lipgloss.Style {
	switch style {
	case Info:
		return ui.button().Copy().Background(ui.Theme.Normal.Cyan).Foreground(ui.Theme.Primary.Background).SetString(body)
	case Warn:
		return ui.button().Copy().Background(ui.Theme.Bright.Yellow).Foreground(ui.Theme.Primary.Background).SetString(body)
	case Error:
		return ui.button().Copy().Background(ui.Theme.Bright.Yellow).Foreground(ui.Theme.Primary.Background).SetString(body)
	case Success:
		return ui.button().Copy().Background(ui.Theme.Bright.Green).Foreground(ui.Theme.Primary.Foreground).SetString(body)
	default:
		return ui.button().SetString(body)
	}
}

func (ui *UI) button() lipgloss.Style {
	return lipgloss.NewStyle().
		Padding(0, 1).
		MarginTop(1)
}
