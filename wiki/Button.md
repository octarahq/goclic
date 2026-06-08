# Button

The `Button` component represents a simple, clickable button in your CLI interface. It is used to trigger a specific action when the user selects it and presses `Enter` or `Space`.

## Usage

```go
package main

import (
    "fmt"
    menu "github.com/octarahq/goclic"
    "github.com/octarahq/goclic/components"
)

func main() {
    m := menu.NewMenu()

    m.Add(components.NewButton("Submit", components.WithButtonOnClick(func() {
        // Action to perform on click
        fmt.Println("Button clicked!")
    })))

    m.Start()
}
```

## Constructor

### `NewButton(label string, opts ...ButtonOption) *Button`

- `label`: The text displayed on the button.
- `opts`: Optional configuration functions (functional options pattern).

### Available Options

- `WithButtonOnClick(func())`: Sets the callback function that is executed when the button is activated.

## Behavior

- **Focus**: Displays `> [Label]` when focused and `  [Label]` otherwise.
- **Activation**: Triggered by pressing `Enter` or `Space`.
