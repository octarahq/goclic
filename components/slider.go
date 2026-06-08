package components

import (
	"fmt"
	"strings"
)

type Slider struct {
	label    string
	value    int
	minValue int
	maxValue int
	step     int
	onChange func(idx int)
}

func NewSlider(label string, value *int, minValue int, maxValue int, step int) *Slider {
	return &Slider{
		label:    label,
		value:    *value,
		minValue: minValue,
		maxValue: maxValue,
		step:     step,
	}
}

func NewSliderC(label string, value *int, minValue int, maxValue int, step int, onChange func(idx int)) *Slider {
	return &Slider{
		label:    label,
		value:    *value,
		minValue: minValue,
		maxValue: maxValue,
		step:     step,
		onChange: onChange,
	}
}

func (d *Slider) GetName() string {
	return "slider"
}

func (d *Slider) Height() int {
	return 1
}

func (d *Slider) IsSelectable() bool {
	return true
}

func (l *Slider) Render(focused bool) string {
	width := 10
	rng := l.maxValue - l.minValue
	filled := 0

	if rng > 0 {
		filled = (l.value - l.minValue) * width / rng
	}

	if filled < 0 {
		filled = 0
	}
	if filled > width {
		filled = width
	}

	slider := fmt.Sprintf("[%s%s]",
		strings.Repeat("█", filled),
		strings.Repeat(" ", width-filled),
	)

	if focused {
		return fmt.Sprintf("> %s : %s %d", l.label, slider, l.value)
	}
	return fmt.Sprintf("  %s : %s %d", l.label, slider, l.value)
}

func (l *Slider) HandleInput(key []byte) bool {
	if len(key) > 0 && key[0] == 27 && len(key) > 2 {
		if key[2] == 'C' {
			if l.value+l.step <= l.maxValue {
				l.value += l.step
			}
			if l.onChange != nil {
				l.onChange(l.value)
			}
			return true
		} else if key[2] == 'D' {
			if l.value-l.step >= l.minValue {
				l.value -= l.step
			}
			if l.onChange != nil {
				l.onChange(l.value)
			}
			return true
		}
	}
	return false
}
