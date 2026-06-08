# Checkbox

The `Checkbox` component allows the user to make a single selection among multiple choices, functioning similarly to a traditional radio button group.

## Usage

```go
package main

import (
    menu "github.com/octarahq/goclic"
    "github.com/octarahq/goclic/components"
)

func main() {
    m := menu.NewMenu()

    var selectedOption int

    m.Add(components.NewCheckbox(
        "Select Environment", 
        []string{"Development", "Staging", "Production"}, 
        &selectedOption,
    ))

    m.Start()
}
```

## Constructor

### `NewCheckbox(label string, options []string, selected *int) *Checkbox`

- `label`: The title displayed above the checkbox options.
- `options`: A slice of strings representing the available choices.
- `selected`: A pointer to an integer that will store the index of the currently selected option.

## Behavior

- **Navigation**: Uses the `Up` and `Down` arrow keys to navigate between the available options.
- **Selection**: Pressing `Enter` or `Space` selects the currently highlighted option and updates the underlying `selected` pointer. Only one option can be selected at a time.
