package main

import (
	"fmt"
	"os"

	menu "github.com/octarahq/goclic"
	"github.com/octarahq/goclic/components"
)

func main() {
	m := menu.NewMenu()

	var darkMode bool
	var volume int = 5
	var difficultyIdx int
	difficulties := []string{"Easy", "Normal", "Hard", "Expert"}

	m.Add(components.NewDisplay("Game Settings", components.WithDisplayBorder(true)))

	m.Add(components.NewSwitch("Dark Mode", &darkMode, components.WithSwitchOnChange(func(val bool) {
		// Logic when toggled
	})))

	m.Add(components.NewSlider("Volume", &volume, 0, 10, 1, components.WithSliderOnChange(func(val int) {
		// Logic on volume change
	})))

	m.Add(components.NewSelect("Difficulty", difficulties, &difficultyIdx, components.WithSelectOnChange(func(idx int) {
		// Logic on difficulty change
	})))

	m.Add(components.NewButton("Save and Quit", components.WithButtonOnClick(func() {
		m.Stop()
		fmt.Println("\n--- Settings Saved ---")
		fmt.Printf("Dark Mode: %v\n", darkMode)
		fmt.Printf("Volume: %d\n", volume)
		fmt.Printf("Difficulty: %s\n", difficulties[difficultyIdx])
		os.Exit(0)
	})))

	m.Start()
}
