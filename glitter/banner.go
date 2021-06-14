package glitter

import (
	"strings"

	"github.com/brittonhayes/glitter/style"
	"github.com/charmbracelet/lipgloss"
)

// Banner is a large text window with title text
func (ui *UI) Banner(body string) lipgloss.Style {
	b := strings.Title(body)
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(ui.Theme.Primary.Foreground).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ui.Theme.Dim.White).
		Padding(style.XS.Size(), style.MD.Size()).
		MarginTop(style.XS.Size()).
		MarginBottom(style.XS.Size()).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true).
		SetString(b)
}
