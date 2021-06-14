package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// Theme is a color scheme
// to be used by all ui components
type Theme struct {
	Primary Primary `json:"primary"`
	Normal  Normal  `json:"normal"`
	Bright  Bright  `json:"bright"`
	Dim     Dim     `json:"dim"`
}

// Primary is the window color scheme
// containing the main foreground and background
// colors
type Primary struct {
	Background    lipgloss.Color `json:"background"`
	Foreground    lipgloss.Color `json:"foreground"`
	DimForeground lipgloss.Color `json:"dim_foreground"`
}

// Normal is the default color set used for most
// operations
type Normal struct {
	Black   lipgloss.Color `json:"black"`
	Blue    lipgloss.Color `json:"blue"`
	Cyan    lipgloss.Color `json:"cyan"`
	Green   lipgloss.Color `json:"green"`
	Magenta lipgloss.Color `json:"magenta"`
	Red     lipgloss.Color `json:"red"`
	White   lipgloss.Color `json:"white"`
	Yellow  lipgloss.Color `json:"yellow"`
}

// Bright is the lighter color set used for
// elements you want highlighted
type Bright struct {
	Black   lipgloss.Color `json:"black"`
	Blue    lipgloss.Color `json:"blue"`
	Cyan    lipgloss.Color `json:"cyan"`
	Green   lipgloss.Color `json:"green"`
	Magenta lipgloss.Color `json:"magenta"`
	Red     lipgloss.Color `json:"red"`
	White   lipgloss.Color `json:"white"`
	Yellow  lipgloss.Color `json:"yellow"`
}

// Dim is the faint color set used for
// elements you'd prefer to be subtle
type Dim struct {
	Black   lipgloss.Color `json:"black"`
	Blue    lipgloss.Color `json:"blue"`
	Cyan    lipgloss.Color `json:"cyan"`
	Green   lipgloss.Color `json:"green"`
	Magenta lipgloss.Color `json:"magenta"`
	Red     lipgloss.Color `json:"red"`
	White   lipgloss.Color `json:"white"`
	Yellow  lipgloss.Color `json:"yellow"`
}
