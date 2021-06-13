package glitter

import (
	"strings"

	"github.com/brittonhayes/glitter/style"
	"github.com/brittonhayes/glitter/theme"
	"github.com/charmbracelet/lipgloss"
)

type UI struct {
	Theme theme.Theme
}

func NewUI(theme theme.Theme) *UI {
	return &UI{Theme: theme}
}

func (ui *UI) Block(s string, options ...lipgloss.WhitespaceOption) string {
	return lipgloss.PlaceHorizontal(80, lipgloss.Center, s, options...)
}

func (ui *UI) Banner(body string) lipgloss.Style {
	b := strings.Title(body)
	return lipgloss.NewStyle().
		Bold(true).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ui.Theme.Dim.White).
		Padding(1, 5).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true).
		SetString(b)
}

func (ui *UI) Border(body string) lipgloss.Style {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(ui.Theme.Primary.Foreground).
		SetString(body)
}

func (ui *UI) Italicize(body string) lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Italic(true).
		Foreground(ui.Theme.Primary.Foreground).
		Background(ui.Theme.Primary.Foreground).
		PaddingLeft(style.PaddingMD).
		PaddingRight(style.PaddingMD).
		Align(lipgloss.Center).
		SetString(body)
}
