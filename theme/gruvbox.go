package theme

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	gruvboxBG      = lipgloss.Color("#282828")
	gruvboxFG      = lipgloss.Color("#EBDBB2")
	gruvboxRed     = lipgloss.Color("#CC241D")
	gruvboxGreen   = lipgloss.Color("#98971A")
	gruvboxBlue    = lipgloss.Color("#458588")
	gruvboxYellow  = lipgloss.Color("#ebcb8b")
	gruvboxMagenta = lipgloss.Color("#B16286")
	gruvboxCyan    = lipgloss.Color("#689D6A")
	gruvboxWhite   = lipgloss.Color("#A89984")
)

// Gruvbox is a lipgloss mapping of the
// Gruvbox color scheme
var Gruvbox = Theme{
	Primary: Primary{
		Background:    gruvboxBG,
		Foreground:    gruvboxFG,
		DimForeground: lipgloss.Color("#928374"),
	},
	Normal: Normal{
		Black:  gruvboxBG,
		Blue:   gruvboxBlue,
		Cyan:   gruvboxCyan,
		Green:  gruvboxGreen,
		Red:    gruvboxRed,
		White:  gruvboxWhite,
		Yellow: gruvboxYellow,
	},
	Bright: Bright{
		Black:   lipgloss.Color("#282828"),
		Blue:    lipgloss.Color("#83A598"),
		Cyan:    lipgloss.Color("#8EC07C"),
		Green:   gruvboxGreen,
		Magenta: gruvboxMagenta,
		Red:     gruvboxRed,
		White:   lipgloss.Color("#EBDBB2"),
		Yellow:  gruvboxYellow,
	},
	Dim: Dim{
		Black:   lipgloss.Color("#373e4d"),
		Blue:    lipgloss.Color("#68809a"),
		Cyan:    lipgloss.Color("#6d96a5"),
		Green:   lipgloss.Color("#809575"),
		Magenta: lipgloss.Color("#8c738c"),
		Red:     lipgloss.Color("#94545d"),
		White:   lipgloss.Color("#928374"),
		Yellow:  lipgloss.Color("#b29e75"),
	},
}
