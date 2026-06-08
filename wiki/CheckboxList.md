# CheckboxList

The `CheckboxList` component provides a list of options where the user can select multiple choices using checkboxes.

## Usage

```go
package main

import (
    menu "github.com/octarahq/goclic"
    "github.com/octarahq/goclic/components"
)

func main() {
    m := menu.NewMenu()

    var checkedItems []bool

    m.Add(components.NewCheckboxList(
        "Select Features", 
        []string{"Feature A", "Feature B", "Feature C"}, 
        &checkedItems,
    ))

    m.Start()
}
```

## Constructor

### `NewCheckboxList(label string, options []string, checked *[]bool) *CheckboxList`

- `label`: The title displayed above the list.
- `options`: A slice of strings representing the available options.
- `checked`: A pointer to a boolean slice that will store the state (checked/unchecked) of each option. Its length will automatically match the `options` length.

## Behavior

- **Navigation**: Uses the `Up` and `Down` arrow keys to navigate between the options in the list.
- **Toggle**: Pressing `Space` or `Enter` toggles the checked state of the currently highlighted option, represented by `[x]` or `[ ]`.
