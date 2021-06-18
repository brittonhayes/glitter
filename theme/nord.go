package theme

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	nordBG     = lipgloss.Color("#2e3440")
	nordWhite  = lipgloss.Color("#eceff4")
	nordDimFG  = lipgloss.Color("#d8dee9")
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
