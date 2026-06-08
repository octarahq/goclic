package components

import (
	"fmt"

	menu "github.com/octarahq/goclic"
)

type Input struct {
	label    string
	value    string
	onChange func(value string)
}

func NewInput(label string, value *string) *Input {
	return &Input{
		label: label,
		value: *value,
	}
}

func NewInputC(label string, value *string, onChange func(string)) *Input {
	return &Input{
		label:    label,
		value:    *value,
		onChange: onChange,
	}
}

func (d *Input) GetName() string {
	return "input"
}

func (d *Input) IsSelectable() bool {
	return true
}

func (l *Input) Render(focused bool) string {
	if focused {
		return fmt.Sprintf("> %s [%s]", l.label, l.value)
	}
	return fmt.Sprintf("  %s [%s]", l.label, l.value)
}

func (l *Input) HandleInput(key []byte) bool {
	if len(key) == 1 {
		b := key[0]

		if b == 8 || b == 127 {
			if len(l.value) > 0 {
				l.value = l.value[:len(l.value)-1]
				if l.onChange != nil {
					l.onChange(l.value)
				}
			}
			return true
		}

		if b == menu.KeyEnter || b == menu.KeySpace {
			return true
		}

		if menu.IsPrintable(key) {
			l.value += string(b)
			if l.onChange != nil {
				l.onChange(l.value)
			}
			return true
		}
	}
	return false
}
