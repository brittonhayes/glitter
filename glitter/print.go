package glitter

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func String(body string, style lipgloss.Style) string {
	return style.Copy().Render(body)
}

func Println(body string, style lipgloss.Style) {
	fmt.Println(style.Copy().Render(body))
}

func Printf(format string, body string, style lipgloss.Style) {
	fmt.Printf(format, style.Copy().Render(body))
}
