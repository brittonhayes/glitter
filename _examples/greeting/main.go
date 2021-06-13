package main

import (
	"fmt"

	"github.com/brittonhayes/glitter/glitter"
	"github.com/brittonhayes/glitter/theme"
)

func main() {
	ui := glitter.NewUI(theme.Gruvbox)

	var (
		banner         = ui.Banner("Welcome to Glitter!")
		comment        = ui.Comment("#", "Enter following command:")
		prefix, prompt = ui.Prompt("$", "go get github.com/brittonhayes/glitter", false)
	)

	// Print the banner
	fmt.Println(banner.String())

	// Print the code comment
	fmt.Println(comment.String())

	// Print the shell prompt
	fmt.Println(prefix.String(), prompt.String())
}
