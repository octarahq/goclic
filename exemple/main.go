package main

import (
	menu "github.com/octarahq/goclic"
	"github.com/octarahq/goclic/components"
)

func main() {
	menu := menu.NewMenu()

	menu.Add(components.NewDisplay("test"))

	menu.Start()
}
