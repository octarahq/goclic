package components

import (
	"fmt"
	"strings"
)

type Display struct {
	label      string
	showBorder bool
}

type DisplayOption func(*Display)

func WithDisplayBorder(showBorder bool) DisplayOption {
	return func(d *Display) {
		d.showBorder = showBorder
	}
}

func NewDisplay(label string, opts ...DisplayOption) *Display {
	d := &Display{
		label:      label,
		showBorder: false,
	}
	for _, opt := range opts {
		opt(d)
	}
	return d
}

func (d *Display) GetName() string {
	return "display"
}

func (d *Display) Height() int {
	h := 1 + strings.Count(d.label, "\n")
	if d.showBorder {
		h += 2 // top and bottom borders
	}
	return h
}

func (d *Display) IsSelectable() bool {
	return false
}

func (l *Display) Render(focused bool) string {
	lines := strings.Split(l.label, "\n")
	var result []string

	if !l.showBorder {
		for _, line := range lines {
			result = append(result, fmt.Sprintf("%s", line))
		}
		return strings.Join(result, "\n")
	}

	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	borderLine := "+" + strings.Repeat("-", maxLen+2) + "+"
	result = append(result, borderLine)
	for _, line := range lines {
		result = append(result, fmt.Sprintf("| %-*s |", maxLen, line))
	}
	result = append(result, borderLine)

	return strings.Join(result, "\n")
}
