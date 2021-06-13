package theme

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	monokaiBG     = lipgloss.Color("#2e2e2e")
	monokaiWhite  = lipgloss.Color("#d6d6d6")
	monokaiDimFG  = lipgloss.Color("#797979")
	monokaiRed    = lipgloss.Color("#b05279")
	monokaiGreen  = lipgloss.Color("#a3be8c")
	monokaiBlue   = lipgloss.Color("#6c99bb")
	monokaiPurple = lipgloss.Color("#9e86c8")
	monokaiPink   = lipgloss.Color("#b05279")
	monokaiYellow = lipgloss.Color("#e5b567")
)

var Monokai = Theme{
	Primary: Primary{
		Background:    monokaiBG,
		Foreground:    monokaiWhite,
		DimForeground: monokaiDimFG,
	},
	Normal: Normal{
		Black:  monokaiBG,
		Blue:   monokaiBlue,
		Cyan:   monokaiPurple,
		Green:  monokaiGreen,
		Red:    monokaiRed,
		White:  monokaiWhite,
		Yellow: monokaiYellow,
	},
	Bright: Bright{
		Black:   monokaiBG,
		Blue:    monokaiBlue,
		Cyan:    monokaiPurple,
		Green:   monokaiGreen,
		Magenta: monokaiPink,
		Red:     monokaiRed,
		White:   monokaiWhite,
	},
	Dim: Dim{
		Black:   monokaiBG,
		Red:     monokaiRed,
		Green:   monokaiGreen,
		Yellow:  monokaiYellow,
		Blue:    monokaiBlue,
		Magenta: monokaiPink,
		Cyan:    monokaiPurple,
		White:   monokaiWhite,
	},
}
