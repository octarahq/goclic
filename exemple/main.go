package main

import (
	menu "github.com/octarahq/goclic"
	"github.com/octarahq/goclic/components"
)

func main() {
	menu := menu.NewMenu()

	darkMode := false
	idx := 0
	volume := 0
	var checked []bool
	var idxChecked int

	menu.Add(components.NewDisplay("test"))
	menu.Add(components.NewSwitch("Darkmode", &darkMode))
	menu.Add(components.NewSelect("Dificultée", []string{"Facile", "Moyen", "Difficile"}, &idx))
	menu.Add(components.NewSlider("Volume", &volume, 0, 10, 1))
	menu.Add(components.NewCheckboxList("Test list", []string{"Oui", "Bien sur", "totalement"}, &checked))
	menu.Add(components.NewCheckbox("Test", []string{"Oui", "Bien sur", "totalement"}, &idxChecked))

	menu.Start()
}
