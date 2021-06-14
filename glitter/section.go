package glitter

import (
	"github.com/brittonhayes/glitter/style"
	"github.com/charmbracelet/lipgloss"
)

// Section is an underlined section divider
func (ui *UI) Section(body string) lipgloss.Style {
	return lipgloss.NewStyle().
		Align(lipgloss.Center).
		Bold(true).
		BorderBottom(true).
		BorderForeground(ui.Theme.Primary.DimForeground).
		BorderStyle(style.BorderDivider).
		Faint(true).
		Foreground(ui.Theme.Primary.Foreground).
		MarginBottom(style.XS.Size()).
		MarginTop(style.XS.Size()).
		MaxWidth(70).
		PaddingLeft(style.XS.Size()).
		PaddingRight(style.XS.Size()).
		Width(30).
		SetString(body)
}

// SectionEnd is an section divider
// with no margins
//
// Use this for ending a section block
func (ui *UI) SectionEnd(body string) lipgloss.Style {
	return lipgloss.NewStyle().
		Align(lipgloss.Center).
		BorderTop(true).
		BorderForeground(ui.Theme.Primary.DimForeground).
		BorderStyle(lipgloss.Border{
			Top:         "_",
			Bottom:      "_",
			Left:        "",
			Right:       "",
			TopLeft:     "",
			TopRight:    "",
			BottomLeft:  "",
			BottomRight: "",
		}).
		Faint(true).
		Foreground(ui.Theme.Primary.Foreground).
		PaddingLeft(style.XS.Size()).
		PaddingRight(style.XS.Size()).
		Width(3).
		SetString(body)
}
