package main

import (
	"fmt"
	"os"

	menu "github.com/octarahq/goclic"
	"github.com/octarahq/goclic/components"
)

func main() {
	m := menu.NewMenu()

	var planIdx int
	plans := []string{"Free", "Pro", "Enterprise"}
	
	var optionsChecked []bool
	options := []string{"24/7 Support", "Auto Backup", "Unlimited Disk Space"}

	m.Add(components.NewDisplay("Subscription Configuration", components.WithDisplayBorder(true)))

	m.Add(components.NewCheckbox("Choose your plan", plans, &planIdx, components.WithCheckboxOnChange(func(idx int) {
		// Logic when plan is chosen
	})))

	m.Add(components.NewDisplay("Additional options (multiple selection):"))

	m.Add(components.NewCheckboxList("Add-ons", options, &optionsChecked, components.WithCheckboxListOnChange(func(idx int) {
		// Logic when checkbox is toggled
	})))

	m.Add(components.NewButton("Confirm", components.WithButtonOnClick(func() {
		m.Stop()
		fmt.Println("\n--- Order Summary ---")
		fmt.Printf("Plan: %s\n", plans[planIdx])
		fmt.Println("Selected Add-ons:")
		for i, checked := range optionsChecked {
			if checked {
				fmt.Printf("- %s\n", options[i])
			}
		}
		os.Exit(0)
	})))

	m.Start()
}
