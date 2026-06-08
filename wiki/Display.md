# Display

The `Display` component is used to output plain text to the console without expecting any user interaction. It's ideal for titles, instructions, or static information.

## Usage

```go
package main

import (
    menu "github.com/octarahq/goclic"
    "github.com/octarahq/goclic/components"
)

func main() {
    m := menu.NewMenu()

    m.Add(components.NewDisplay("--- Main Menu ---"))
    m.Add(components.NewDisplay("Please select an option below:"))

    // Add other interactive components here

    m.Start()
}
```

## Constructor

### `NewDisplay(text string) *Display`

- `text`: The string of text to be displayed.

## Behavior

- **Non-interactive**: The user cannot focus, click, or interact with this component. The cursor will simply skip over it when navigating through the menu.
