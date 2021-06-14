package style

import (
	"github.com/charmbracelet/lipgloss"
)

// Border is a specifier for
// the type of border
// to use on a component
type Border int

const (
	// Default border is a simple
	// text border with straight
	// lines
	Default Border = iota

	// Tab is a border with rounded
	// tops and a straight bottom
	Tab

	// Rounded is a border that
	// has all corners rounded
	Rounded

	// Crossed is a border with
	// X characters on the corners
	Crossed

	// Thick is a border with a bold
	// line
	Thick
)

var (
	// BorderTab is a border with rounded
	// tops and a straight bottom
	BorderTab = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	// BorderCrossed is a border with
	// X characters on the corners
	BorderCrossed = lipgloss.Border{
		Top:         "¯",
		Bottom:      "_",
		Left:        "│",
		Right:       "│",
		TopLeft:     "+",
		TopRight:    "+",
		BottomLeft:  "+",
		BottomRight: "+",
	}

	// BorderDivider is a section divider with
	// lines only on the bottom
	BorderDivider = lipgloss.Border{
		Top:         "",
		Bottom:      ".",
		Left:        "",
		Right:       "",
		TopLeft:     "",
		TopRight:    "",
		BottomLeft:  "",
		BottomRight: "",
	}
)
