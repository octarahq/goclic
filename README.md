# goclic

A Go module for easily creating interactive command-line interfaces (CLI) with ready-to-use components.

## Installation

```bash
go get github.com/octarahq/goclic
```

## Available components

[Full documentation](https://github.com/octarahq/goclic/wiki)

- `Display`: Text display.
- `Switch`: On/Off toggle.
- `Select`: Choose one option from a dropdown list.
- `Slider`: Select a numeric value within a range.
- `CheckboxList`: Multiple selection with checkboxes.
- `Checkbox`: Single selection among multiple choices (radio button style).
- `Input`: Text input field.
- `Button`: Simple button to trigger an action.

## Usage

Example of creating an interactive menu with multiple components:

```go
package main

import (
    menu "github.com/octarahq/goclic"
    "github.com/octarahq/goclic/components"
)

func main() {
    m := menu.NewMenu()

    darkMode := false
    difficultyIdx := 0
    volume := 0
    var checkedList []bool
    var checkedOption int

    m.Add(components.NewDisplay("Main menu"))
    m.Add(components.NewSwitch("Dark mode", &darkMode))
    m.Add(components.NewSelect("Difficulty", []string{"Easy", "Medium", "Hard"}, &difficultyIdx))
    m.Add(components.NewSlider("Volume", &volume, 0, 10, 1))
    m.Add(components.NewCheckboxList("Advanced settings", []string{"Option A", "Option B", "Option C"}, &checkedList))
    m.Add(components.NewCheckbox("Server selection", []string{"Server 1", "Server 2", "Server 3"}, &checkedOption))

    m.Start()
}
```
