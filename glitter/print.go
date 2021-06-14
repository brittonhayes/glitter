package glitter

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// String converts any lipgloss style to
// a string
func String(body string, style lipgloss.Style) string {
	return style.Copy().Render(body)
}

// Println prints a lipgloss style and text body
// to the screen with a new line
func Println(body string, style lipgloss.Style) {
	fmt.Println(style.Copy().Render(body))
}

// Printf prints a lipgloss style and text body
// to the screen with a custom format
func Printf(format string, body string, style lipgloss.Style) {
	fmt.Printf(format, style.Copy().Render(body))
}
