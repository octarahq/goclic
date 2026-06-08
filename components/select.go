package components

import (
	"fmt"
)

type Select struct {
	label       string
	options     []string
	selectedIdx int
	onChange    func(idx int)
}

func NewSelect(label string, options []string, selectedIdx *int) *Select {
	return &Select{
		label:       label,
		options:     options,
		selectedIdx: *selectedIdx,
	}
}

func NewSelectC(label string, options []string, selectedIdx *int, onChange func(idx int)) *Select {
	return &Select{
		label:       label,
		options:     options,
		selectedIdx: *selectedIdx,
		onChange:    onChange,
	}
}

func (d *Select) GetName() string {
	return "select"
}

func (d *Select) IsSelectable() bool {
	return true
}

func (l *Select) Render(focused bool) string {
	if focused {
		return fmt.Sprintf("> %s : < %s >", l.label, l.options[l.selectedIdx])
	}
	return fmt.Sprintf("  %s : %s", l.label, l.options[l.selectedIdx])
}

func (l *Select) HandleInput(key []byte) bool {
	if len(key) > 0 && key[0] == 27 && len(key) > 2 {
		if key[2] == 'C' {
			if l.selectedIdx+1 > len(l.options)-1 {
				l.selectedIdx = 0
			} else {
				l.selectedIdx++
			}
			if l.onChange != nil {
				l.onChange(l.selectedIdx)
			}
			return true
		} else if key[2] == 'D' {
			if l.selectedIdx-1 < 0 {
				l.selectedIdx = len(l.options) - 1
			} else {
				l.selectedIdx--
			}
			if l.onChange != nil {
				l.onChange(l.selectedIdx)
			}
			return true
		}
	}
	return false
}
