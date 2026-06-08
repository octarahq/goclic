# Slider

The `Slider` component provides an interface for users to visually select a numeric value within a defined range.

## Usage

```go
package main

import (
    menu "github.com/octarahq/goclic"
    "github.com/octarahq/goclic/components"
)

func main() {
    m := menu.NewMenu()

    var volume int = 5

    m.Add(components.NewSlider(
        "Volume", 
        &volume, 
        0, 10, 1, // min, max, step
    ))

    m.Start()
}
```

## Constructor

### `NewSlider(label string, value *int, min, max, step int, opts ...SliderOption) *Slider`

- `label`: The text label displayed next to the slider.
- `value`: A pointer to an integer where the current value is stored. This also acts as the initial value when the menu is rendered.
- `min`: The minimum allowed value.
- `max`: The maximum allowed value.
- `step`: The amount the value changes per key press.
- `opts`: Optional configuration functions (functional options pattern).

### Available Options

- `WithSliderOnChange(func(idx int))`: Executes a callback function whenever the slider value changes.

## Behavior

- **Navigation**: While the component is focused, pressing `Left` or `Right` arrow keys decreases or increases the value by the `step` amount, clamped between `min` and `max`.
- **Visuals**: Displays a visual bar `[====----]` representing the current value relative to the range.
