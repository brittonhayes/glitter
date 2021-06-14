package glitter

import (
	"github.com/brittonhayes/glitter/style"
	"github.com/charmbracelet/lipgloss"
)

// Border is a bordered text box with a few different styles available
func (ui *UI) Border(body string, border style.Border) lipgloss.Style {
	base := lipgloss.NewStyle().
		PaddingLeft(style.XS.Size()).
		PaddingRight(style.XS.Size()).
		Foreground(ui.Theme.Primary.Foreground).
		BorderForeground(ui.Theme.Primary.Foreground).
		SetString(body)
	switch border {
	case style.Tab:
		return base.Copy().BorderStyle(style.BorderTab)
	case style.Crossed:
		return base.Copy().BorderStyle(style.BorderCrossed)
	case style.Rounded:
		return base.Copy().BorderStyle(lipgloss.RoundedBorder())
	case style.Thick:
		return base.Copy().BorderStyle(lipgloss.ThickBorder())
	default:
		return base.Copy().BorderStyle(lipgloss.NormalBorder())
	}
}
