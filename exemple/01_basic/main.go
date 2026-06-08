package main

import (
	"fmt"
	"os"

	menu "github.com/octarahq/goclic"
	"github.com/octarahq/goclic/components"
)

func main() {
	m := menu.NewMenu()
	var username string

	m.Add(components.NewDisplay("=== Welcome ===", components.WithDisplayBorder(true)))
	
	m.Add(components.NewInput("What is your name?", &username, components.WithInputOnChange(func(val string) {
		// Optional callback when value changes
	})))

	m.Add(components.NewButton("Submit", components.WithButtonOnClick(func() {
		m.Stop()
		fmt.Printf("\nHello %s! Nice to meet you.\n", username)
		os.Exit(0)
	})))

	m.Start()
}
