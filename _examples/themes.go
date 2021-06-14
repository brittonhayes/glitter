package main

import (
	"fmt"

	"github.com/brittonhayes/glitter/glitter"
	"github.com/brittonhayes/glitter/style"
	"github.com/brittonhayes/glitter/theme"
)

func main() {
	// Create a new glitter UI and select a theme
	// ui := glitter.NewUI(theme.Nord)
	ui := glitter.NewUI(theme.Gruvbox)
	// ui := glitter.NewUI(theme.Monokai)

	// Create a few UI elements
	var (
		btn            = ui.Button("Info", style.Info)
		btnSuccess     = ui.Button("Success", style.Success)
		btnWarn        = ui.Button("Warn", style.Warn)
		btnError       = ui.Button("Error", style.Error)
		italics        = ui.Italicize("Some italicized text")
		borderTab      = ui.Border("Tab Border", style.Tab)
		borderRound    = ui.Border("Rounded Border", style.Rounded)
		borderCrossed  = ui.Border("Bordered Crossed", style.Crossed)
		prefix, prompt = ui.Prompt(">", "echo hello glitter", false)
		banner         = ui.Banner("Welcome to my app!")
	)

	// Render the elements to see
	// which theme you like best
	fmt.Println(ui.Section("Text Styles"))
	fmt.Println(ui.Comment("#", "You can add comments, prompts, italics and more"))
	fmt.Println(prefix, prompt)
	fmt.Println(italics)

	fmt.Println(ui.Section("Buttons"))
	fmt.Println(ui.Comment("#", "Try all the buttons!"))
	fmt.Printf("%s %s %s %s\n", btn, btnSuccess, btnWarn, btnError)

	fmt.Println(ui.Section("Banners"))
	fmt.Println(ui.Comment("#", "Bannders for when you want to make an entrance"))
	fmt.Println(banner)

	fmt.Println(ui.Section("Borders"))
	fmt.Println(ui.Comment("#", "Borders for when you're feeling edgy"))
	fmt.Println(borderTab)
	fmt.Println(borderRound)
	fmt.Println(borderCrossed)
}
