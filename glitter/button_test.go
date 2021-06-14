package glitter_test

import (
	"fmt"

	"github.com/brittonhayes/glitter/glitter"
	"github.com/brittonhayes/glitter/style"
	"github.com/brittonhayes/glitter/theme"
)

func ExampleUI_Button() {
	ui := glitter.NewUI(theme.Gruvbox)
	button := ui.Button("Some button test", style.Info)
	fmt.Println(button.String())
}
