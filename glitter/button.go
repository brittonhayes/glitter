package glitter

import (
	"github.com/brittonhayes/glitter/style"
	"github.com/charmbracelet/lipgloss"
)

// Button is a button shaped UI element with a few different
// style options
func (ui *UI) Button(body string, buttonStyle style.ButtonStyle) lipgloss.Style {
	switch buttonStyle {
	case style.Info:
		return ui.button().Copy().Background(ui.Theme.Normal.Cyan).Foreground(ui.Theme.Primary.Background).SetString(body)
	case style.Warn:
		return ui.button().Copy().Background(ui.Theme.Bright.Yellow).Foreground(ui.Theme.Primary.Background).SetString(body)
	case style.Error:
		return ui.button().Copy().Background(ui.Theme.Bright.Red).Foreground(ui.Theme.Primary.Background).SetString(body)
	case style.Success:
		return ui.button().Copy().Background(ui.Theme.Bright.Green).Foreground(ui.Theme.Primary.Background).SetString(body)
	default:
		return ui.button().SetString(body)
	}
}

func (ui *UI) button() lipgloss.Style {
	return lipgloss.NewStyle().
		MaxWidth(30).
		Padding(0, style.XS.Size())
}
