package components

import (
	"fmt"
	"strings"
)

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

func (d *Display) Height() int {
	return 1 + strings.Count(d.label, "\n")
}

func (d *Display) IsSelectable() bool {
	return false
}

func (l *Display) Render(focused bool) string {
	lines := strings.Split(l.label, "\n")
	var result []string
	for _, line := range lines {
		result = append(result, fmt.Sprintf("| %s |", line))
	}
	return strings.Join(result, "\n")
}
