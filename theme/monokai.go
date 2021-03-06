package theme

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	monokaiBG     = lipgloss.Color("#2e2e2e")
	monokaiWhite  = lipgloss.Color("#d6d6d6")
	monokaiDimFG  = lipgloss.Color("#797979")
	monokaiRed    = lipgloss.Color("#b05279")
	monokaiGreen  = lipgloss.Color("#b4d273")
	monokaiBlue   = lipgloss.Color("#6c99bb")
	monokaiPurple = lipgloss.Color("#9e86c8")
	monokaiPink   = lipgloss.Color("#b05279")
	monokaiYellow = lipgloss.Color("#e5b567")
)

// Monokai is a lipgloss mapping of the
// Monokai color scheme
var Monokai = Theme{
	Primary: Primary{
		Background:    monokaiBG,
		Foreground:    monokaiWhite,
		DimForeground: monokaiDimFG,
	},
	Normal: Normal{
		Black:   monokaiBG,
		Blue:    monokaiBlue,
		Cyan:    monokaiPurple,
		Green:   monokaiGreen,
		Magenta: monokaiPink,
		Red:     monokaiRed,
		White:   monokaiWhite,
		Yellow:  monokaiYellow,
	},
	Bright: Bright{
		Black:   monokaiBG,
		Blue:    monokaiBlue,
		Cyan:    monokaiPurple,
		Green:   monokaiGreen,
		Magenta: monokaiPink,
		Red:     monokaiRed,
		White:   monokaiWhite,
		Yellow:  monokaiYellow,
	},
	Dim: Dim{
		Black:   monokaiBG,
		Blue:    monokaiBlue,
		Cyan:    monokaiPurple,
		Green:   monokaiGreen,
		Magenta: monokaiPink,
		Red:     monokaiRed,
		White:   monokaiWhite,
		Yellow:  monokaiYellow,
	},
}
