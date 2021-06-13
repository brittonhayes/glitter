package main

import (
	"fmt"

	"github.com/brittonhayes/glitter/glitter"
	"github.com/brittonhayes/glitter/theme"
)

func main() {
	// Create a new glitter UI and select a theme
	ui := glitter.NewUI(theme.Monokai)

	// Create a button and mark it as active
	btn := ui.Button("hello there", glitter.Info)

	// Render the button!
	fmt.Println(btn)
}
