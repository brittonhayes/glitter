package glitter

import (
	"github.com/brittonhayes/glitter/theme"
	"github.com/charmbracelet/lipgloss"
)

// UI is the building
// block for all components
//
// The UI maps the provided theme
// to each rendered component
type UI struct {
	Theme theme.Theme
}

// NewUI creates a new instance of the UI
// type with a custom theme provided
func NewUI(theme theme.Theme) *UI {
	return &UI{Theme: theme}
}

// Block is a horizontally aligned text block
func (ui *UI) Block(s string, options ...lipgloss.WhitespaceOption) string {
	return lipgloss.PlaceHorizontal(80, lipgloss.Center, s, options...)
}
