package components

import (
	"fmt"
)

type Select struct {
	label       string
	options     []string
	selectedIdx *int
	onChange    func(idx int)
}

type SelectOption func(*Select)

func WithSelectOnChange(onChange func(idx int)) SelectOption {
	return func(s *Select) {
		s.onChange = onChange
	}
}

func NewSelect(label string, options []string, selectedIdx *int, opts ...SelectOption) *Select {
	s := &Select{
		label:       label,
		options:     options,
		selectedIdx: selectedIdx,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (d *Select) GetName() string {
	return "select"
}

func (d *Select) Height() int {
	return 1
}

func (d *Select) IsSelectable() bool {
	return true
}

func (l *Select) Render(focused bool) string {
	idx := 0
	if l.selectedIdx != nil {
		idx = *l.selectedIdx
	}
	if focused {
		return fmt.Sprintf("> %s : < %s >", l.label, l.options[idx])
	}
	return fmt.Sprintf("  %s : %s", l.label, l.options[idx])
}

func (l *Select) HandleInput(key []byte) bool {
	if len(key) > 0 && key[0] == 27 && len(key) > 2 {
		if key[2] == 'C' {
			if l.selectedIdx != nil {
				if *l.selectedIdx+1 > len(l.options)-1 {
					*l.selectedIdx = 0
				} else {
					*l.selectedIdx++
				}
				if l.onChange != nil {
					l.onChange(*l.selectedIdx)
				}
			}
			return true
		} else if key[2] == 'D' {
			if l.selectedIdx != nil {
				if *l.selectedIdx-1 < 0 {
					*l.selectedIdx = len(l.options) - 1
				} else {
					*l.selectedIdx--
				}
				if l.onChange != nil {
					l.onChange(*l.selectedIdx)
				}
			}
			return true
		}
	}
	return false
}
