package components

import (
	"fmt"
	"strings"
)

type Slider struct {
	label    string
	value    *int
	minValue int
	maxValue int
	step     int
	onChange func(idx int)
}

type SliderOption func(*Slider)

func WithSliderOnChange(onChange func(idx int)) SliderOption {
	return func(s *Slider) {
		s.onChange = onChange
	}
}

func NewSlider(label string, value *int, minValue int, maxValue int, step int, opts ...SliderOption) *Slider {
	s := &Slider{
		label:    label,
		value:    value,
		minValue: minValue,
		maxValue: maxValue,
		step:     step,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
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
	val := 0

	if l.value != nil {
		val = *l.value
	}

	if rng > 0 {
		filled = (val - l.minValue) * width / rng
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
		return fmt.Sprintf("> %s : %s %d", l.label, slider, val)
	}
	return fmt.Sprintf("  %s : %s %d", l.label, slider, val)
}

func (l *Slider) HandleInput(key []byte) bool {
	if len(key) > 0 && key[0] == 27 && len(key) > 2 {
		if key[2] == 'C' {
			if l.value != nil && *l.value+l.step <= l.maxValue {
				*l.value += l.step
			}
			if l.onChange != nil && l.value != nil {
				l.onChange(*l.value)
			}
			return true
		} else if key[2] == 'D' {
			if l.value != nil && *l.value-l.step >= l.minValue {
				*l.value -= l.step
			}
			if l.onChange != nil && l.value != nil {
				l.onChange(*l.value)
			}
			return true
		}
	}
	return false
}
