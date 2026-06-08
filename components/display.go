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

func (d *Display) IsSelectable() bool {
	return false
}

func (l *Display) Render(focused bool) string {
	return fmt.Sprintf("  | %s |", l.label)
}
