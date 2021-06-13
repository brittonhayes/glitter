package glitter

import (
	"github.com/charmbracelet/lipgloss"
)

func (ui *UI) Button(body string, active bool) string {
	if active {
		return ui.buttonActive().Render(body)
	}
	return ui.button().Render(body)
}

func (ui *UI) button() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Foreground(ui.Theme.Primary.Foreground).
		Background(ui.Theme.Primary.Background).
		Padding(0, 1).
		MarginTop(1)
}

func (ui *UI) buttonActive() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(ui.Theme.Primary.Background).
		Background(ui.Theme.Primary.Foreground).
		Padding(0, 1).
		Margin(1)
}
