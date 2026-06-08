# Switch

The `Switch` component acts as a simple boolean toggle, commonly used to represent On/Off or True/False states.

## Usage

```go
package main

import (
    menu "github.com/octarahq/goclic"
    "github.com/octarahq/goclic/components"
)

func main() {
    m := menu.NewMenu()

    var enableFeature bool

    m.Add(components.NewSwitch("Enable Feature X", &enableFeature))

    m.Start()
}
```

## Constructor

### `NewSwitch(label string, value *bool) *Switch`

- `label`: The text description of the toggle.
- `value`: A pointer to a boolean representing the state. Pass `true` initially if you want the switch to default to On.

## Behavior

- **Toggle**: While the component is focused, pressing `Enter`, `Space`, `Left`, or `Right` toggles the boolean value between `true` and `false`.
- **Visuals**: Displays `[ON]` when true and `[OFF]` when false.
