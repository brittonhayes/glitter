package style

// ButtonStyle maps
// each button aesthetic to
// an integer
type ButtonStyle int

const (
	// Info is a button with a
	// primary foreground color
	Info ButtonStyle = iota

	// Warn is a button with a
	// yellow to indicate an issue
	Warn

	// Error is a button with a
	// red color to indicate a critical
	// problem
	Error

	// Success is a button with a
	// green color to indicate
	// everything is good
	Success
)

