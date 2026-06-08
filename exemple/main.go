package main

import (
	menu "github.com/octarahq/goclic"
	"github.com/octarahq/goclic/components"
)

func main() {
	menu := menu.NewMenu()

	darkMode := false

	menu.Add(components.NewDisplay("test"))
	menu.Add(components.NewSwitch("Darkmode", &darkMode))

	menu.Start()
}
