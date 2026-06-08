package components

import (
	"fmt"
	menu "github.com/octarahq/goclic"
)

type Switch struct {
	label    string
	value    *bool
	onChange func(value bool)
}

func NewSwitch(label string, value *bool) *Switch {
	return &Switch{
		label: label,
		value: value,
	}
}

func NewSwitchC(label string, value *bool, onChange func(bool)) *Switch {
	return &Switch{
		label:    label,
		value:    value,
		onChange: onChange,
	}
}

func (d *Switch) GetName() string {
	return "switch"
}

func (d *Switch) IsSelectable() bool {
	return true
}

func (l *Switch) Render(focused bool) string {
	value := " "
	if *l.value {
		value = "X"
	}

	if focused {
		return fmt.Sprintf("> [%s] %s", value, l.label)
	}
	return fmt.Sprintf("  [%s] %s", value, l.label)
}

func (l *Switch) HandleInput(key []byte) bool {
	if len(key) == 1 && (key[0] == menu.KeyEnter || key[0] == menu.KeySpace) {
		*l.value = !*l.value

		if l.onChange != nil {
			l.onChange(*l.value)
		}
		return true
	}
	return false
}
