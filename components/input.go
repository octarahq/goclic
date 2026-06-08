package components

import (
	"fmt"

	"github.com/atotto/clipboard"
	menu "github.com/octarahq/goclic"
)

type Input struct {
	label    string
	value    *string
	onChange func(value string)
}

func NewInput(label string, value *string) *Input {
	return &Input{
		label: label,
		value: value,
	}
}

func NewInputC(label string, value *string, onChange func(string)) *Input {
	return &Input{
		label:    label,
		value:    value,
		onChange: onChange,
	}
}

func (d *Input) GetName() string {
	return "input"
}

func (d *Input) Height() int {
	return 1
}

func (d *Input) IsSelectable() bool {
	return true
}

func (l *Input) Render(focused bool) string {
	val := ""
	if l.value != nil {
		val = *l.value
	}
	if focused {
		return fmt.Sprintf("> %s [%s]", l.label, val)
	}
	return fmt.Sprintf("  %s [%s]", l.label, val)
}

func (l *Input) HandleInput(key []byte) bool {
	if len(key) == 0 {
		return false
	}

	if len(key) == 1 && key[0] == menu.KeyCtrlV {
		text, err := clipboard.ReadAll()
		if err == nil && len(text) > 0 {
			*l.value += text
			if l.onChange != nil {
				l.onChange(*l.value)
			}
			return true
		}
	}

	if len(key) > 1 && menu.IsPrintable(key) {
		*l.value += string(key)
		if l.onChange != nil {
			l.onChange(*l.value)
		}
		return true
	}

	if len(key) == 1 {
		b := key[0]

		if b == 8 || b == 127 {
			if len(*l.value) > 0 {
				*l.value = (*l.value)[:len(*l.value)-1]
				if l.onChange != nil {
					l.onChange(*l.value)
				}
			}
			return true
		}

		if b == menu.KeyEnter {
			return true
		}

		if b == menu.KeySpace {
			*l.value += " "
			if l.onChange != nil {
				l.onChange(*l.value)
			}
			return true
		}

		if menu.IsPrintable(key) {
			*l.value += string(b)
			if l.onChange != nil {
				l.onChange(*l.value)
			}
			return true
		}
	}
	return false
}
