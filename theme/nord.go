package theme

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	nordBG     = lipgloss.Color("#2e3440")
	nordWhite  = lipgloss.Color("#e5e9f0")
	nordDimFG  = lipgloss.Color("#4c566a")
	nordRed    = lipgloss.Color("#bf616a")
	nordGreen  = lipgloss.Color("#a3be8c")
	nordBlue   = lipgloss.Color("#5e81ac")
	nordCyan   = lipgloss.Color("#88c0d0")
	nordPink   = lipgloss.Color("#b48ead")
	nordYellow = lipgloss.Color("#ebcb8b")
)

// Nord is a lipgloss mapping of the
// Nord color scheme
var Nord = Theme{
	Primary: Primary{
		Background:    nordBG,
		Foreground:    nordWhite,
		DimForeground: nordDimFG,
	},
	Normal: Normal{
		Black:   nordBG,
		Blue:    nordBlue,
		Cyan:    nordCyan,
		Green:   nordGreen,
		Magenta: nordPink,
		Red:     nordRed,
		White:   nordWhite,
		Yellow:  nordYellow,
	},
	Bright: Bright{
		Black:   nordBG,
		Blue:    nordBlue,
		Cyan:    nordCyan,
		Green:   nordGreen,
		Magenta: nordPink,
		Red:     nordRed,
		White:   nordWhite,
		Yellow:  nordYellow,
	},
	Dim: Dim{
		Black:   nordBG,
		Blue:    nordBlue,
		Cyan:    nordCyan,
		Green:   nordGreen,
		Magenta: nordPink,
		Red:     nordRed,
		White:   nordDimFG,
		Yellow:  nordYellow,
	},
}
