# Input

The `Input` component allows the user to type a free-form text string directly into the CLI.

## Usage

```go
package main

import (
    menu "github.com/octarahq/goclic"
    "github.com/octarahq/goclic/components"
)

func main() {
    m := menu.NewMenu()

    var username string

    m.Add(components.NewInput("Username", &username))

    m.Start()
}
```

## Constructor

### `NewInput(label string, value *string) *Input`

- `label`: The prompt or label displayed next to the input field.
- `value`: A pointer to a string that will store the user's typed input.

## Behavior

- **Focus**: When focused, the user can start typing. The text is immediately reflected on screen and stored in the underlying string pointer.
- **Editing**: Supports `Backspace` for correcting text.
- **Exit**: The user can press `Up`, `Down`, or `Enter` to finish editing and navigate away from the input field.
