package theme

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	gruvboxBG      = lipgloss.Color("#2e3440")
	gruvboxFG      = lipgloss.Color("#d8dee9")
	gruvboxRed     = lipgloss.Color("#bf616a")
	gruvboxGreen   = lipgloss.Color("#a3be8c")
	gruvboxBlue    = lipgloss.Color("#81a1c1")
	gruvboxYellow  = lipgloss.Color("#ebcb8b")
	gruvboxMagenta = lipgloss.Color("#b48ead")
)

var Gruvbox = Theme{
	Primary: Primary{
		Background:    gruvboxBG,
		Foreground:    gruvboxFG,
		DimForeground: lipgloss.Color("#a5abb6"),
	},
	Cursor: Cursor{
		Text:   gruvboxBG,
		Cursor: gruvboxFG,
	},
	Selection: Selection{
		Background: lipgloss.Color("#4c566a"),
	},
	Normal: Normal{
		Black:  lipgloss.Color("#3b4252"),
		Blue:   gruvboxBlue,
		Cyan:   lipgloss.Color("#88c0d0"),
		Green:  gruvboxGreen,
		Red:    gruvboxRed,
		White:  lipgloss.Color("#e5e9f0"),
		Yellow: gruvboxYellow,
	},
	Bright: Bright{
		Black:   lipgloss.Color("#4c566a"),
		Blue:    gruvboxBlue,
		Cyan:    lipgloss.Color("#8fbcbb"),
		Green:   gruvboxGreen,
		Magenta: gruvboxMagenta,
		Red:     gruvboxRed,
		White:   lipgloss.Color("#eceff4"),
	},
	Dim: Dim{
		Black:   lipgloss.Color("#373e4d"),
		Red:     lipgloss.Color("#94545d"),
		Green:   lipgloss.Color("#809575"),
		Yellow:  lipgloss.Color("#b29e75"),
		Blue:    lipgloss.Color("#68809a"),
		Magenta: lipgloss.Color("#8c738c"),
		Cyan:    lipgloss.Color("#6d96a5"),
		White:   lipgloss.Color("#aeb3bb"),
	},
}
