package theme

import (
	"github.com/charmbracelet/lipgloss"
)

type Theme struct {
	Primary Primary `json:"primary"`
	Normal  Normal  `json:"normal"`
	Bright  Bright  `json:"bright"`
	Dim     Dim     `json:"dim"`
}

type Primary struct {
	Background    lipgloss.Color `json:"background"`
	Foreground    lipgloss.Color `json:"foreground"`
	DimForeground lipgloss.Color `json:"dim_foreground"`
}

type Normal struct {
	Black   lipgloss.Color `json:"black"`
	Red     lipgloss.Color `json:"red"`
	Green   lipgloss.Color `json:"green"`
	Yellow  lipgloss.Color `json:"yellow"`
	Blue    lipgloss.Color `json:"blue"`
	Magenta lipgloss.Color `json:"magenta"`
	Cyan    lipgloss.Color `json:"cyan"`
	White   lipgloss.Color `json:"white"`
}

type Bright struct {
	Black   lipgloss.Color `json:"black"`
	Red     lipgloss.Color `json:"red"`
	Green   lipgloss.Color `json:"green"`
	Yellow  lipgloss.Color `json:"yellow"`
	Blue    lipgloss.Color `json:"blue"`
	Magenta lipgloss.Color `json:"magenta"`
	Cyan    lipgloss.Color `json:"cyan"`
	White   lipgloss.Color `json:"white"`
}

type Dim struct {
	Black   lipgloss.Color `json:"black"`
	Red     lipgloss.Color `json:"red"`
	Green   lipgloss.Color `json:"green"`
	Yellow  lipgloss.Color `json:"yellow"`
	Blue    lipgloss.Color `json:"blue"`
	Magenta lipgloss.Color `json:"magenta"`
	Cyan    lipgloss.Color `json:"cyan"`
	White   lipgloss.Color `json:"white"`
}
