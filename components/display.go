package components

import "fmt"

type Display struct {
	label string
}

func NewDisplay(label string) *Display {
	return &Display{
		label: label,
	}
}

func (d *Display) GetName() string {
	return "display"
}

func (l *Display) Render(focused bool) string {
	if focused {
		return fmt.Sprintf("> | %s |", l.label)
	} else {
		return fmt.Sprintf("  | %s |", l.label)
	}
}
