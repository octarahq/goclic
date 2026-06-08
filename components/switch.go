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

type SwitchOption func(*Switch)

func WithSwitchOnChange(onChange func(bool)) SwitchOption {
	return func(s *Switch) {
		s.onChange = onChange
	}
}

func NewSwitch(label string, value *bool, opts ...SwitchOption) *Switch {
	s := &Switch{
		label: label,
		value: value,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (d *Switch) GetName() string {
	return "switch"
}

func (d *Switch) Height() int {
	return 1
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
