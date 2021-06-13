package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/brittonhayes/pkg/temple"
)

var (
	//go:embed readme.tmpl
	readme string
)

func main() {
	simple, err := os.ReadFile("_examples/simple/button.go")
	if err != nil {
		log.Fatalln(err)
	}

	data := map[string]interface{}{
		"Install": "go get github.com/brittonhayes/glitter",
		"Example": string(simple),
	}

	tpl, err := temple.Render(readme, &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(tpl.String())
}
