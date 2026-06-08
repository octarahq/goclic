# Select

The `Select` component displays a dropdown-style list where the user can pick exactly one option from the provided choices.

## Usage

```go
package main

import (
    menu "github.com/octarahq/goclic"
    "github.com/octarahq/goclic/components"
)

func main() {
    m := menu.NewMenu()

    var selectedIndex int

    m.Add(components.NewSelect(
        "Difficulty Level", 
        []string{"Easy", "Medium", "Hard"}, 
        &selectedIndex,
    ))

    m.Start()
}
```

## Constructor

### `NewSelect(label string, options []string, selected *int, opts ...SelectOption) *Select`

- `label`: The title displayed next to the currently selected option.
- `options`: A slice of strings representing the available choices in the dropdown.
- `selected`: A pointer to an integer that stores the index of the currently selected option.
- `opts`: Optional configuration functions (functional options pattern).

### Available Options

- `WithSelectOnChange(func(idx int))`: Executes a callback function whenever the selected option changes.

## Behavior

- **Navigation**: While the component is focused, pressing `Left` or `Right` arrow keys cycles through the available options.
- **Update**: The `selected` integer is updated live as the user navigates through the choices.
