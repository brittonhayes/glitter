package theme

import "github.com/charmbracelet/lipgloss"

const (
	oneFG = lipgloss.Color("#f0f6f0")
	oneBG = lipgloss.Color("#222323")
)

var OneBit = Theme{
	Primary: Primary{
		Foreground: oneFG,
		Background: oneBG,
	},
	Normal: Normal{
		Black:   oneBG,
		Blue:    oneFG,
		Cyan:    oneFG,
		Green:   oneFG,
		Magenta: oneFG,
		Red:     oneFG,
		White:   oneFG,
		Yellow:  oneFG,
	},
	Bright: Bright{
		Black:   oneBG,
		Blue:    oneFG,
		Cyan:    oneFG,
		Green:   oneFG,
		Magenta: oneFG,
		Red:     oneFG,
		White:   oneFG,
		Yellow:  oneFG,
	},
	Dim: Dim{
		Black:   oneBG,
		Blue:    oneFG,
		Cyan:    oneFG,
		Green:   oneFG,
		Magenta: oneFG,
		Red:     oneFG,
		White:   oneFG,
		Yellow:  oneFG,
	},
}
