package style

// Size is a specifier for
// how large or how the elements
// of a component should be
type Size int

const (
	// XS is the smallest available
	// size
	XS Size = iota + 1
	_

	// SM is a small element size
	SM
	_

	// MD is a the standard element size of
	// medium
	MD
	_

	// LG is a bigger element size used for
	// titles
	LG
	_
	_

	// XL is a very large element size used
	// for big titles and banners
	XL
	_
	_
	_
	_

	// XXL is a massive element size for when
	// you want your component to dominate the
	// screen
	XXL
)

// Size converts the available
// sizes to an integer for use
// with lipgloss
func (s Size) Size() int {
	return int(s)
}
