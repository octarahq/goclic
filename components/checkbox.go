package components

import (
	"fmt"
	"strings"

	menu "github.com/octarahq/goclic"
)

type Checkbox struct {
	label       string
	options     []string
	checked     *int
	internalIdx int
	onChange    func(idx int)
}

type CheckboxOption func(*Checkbox)

func WithCheckboxOnChange(onChange func(idx int)) CheckboxOption {
	return func(c *Checkbox) {
		c.onChange = onChange
	}
}

func NewCheckbox(label string, options []string, checked *int, opts ...CheckboxOption) *Checkbox {
	c := &Checkbox{
		label:       label,
		options:     options,
		checked:     checked,
		internalIdx: 0,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Checkbox) GetName() string    { return "checkbox" }
func (c *Checkbox) IsSelectable() bool { return true }

func (c *Checkbox) Height() int {
	return 1 + len(c.options)
}

func (c *Checkbox) Render(focused bool) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("  %s :\r\n", c.label))

	for i, option := range c.options {
		box := " "
		if c.checked != nil && i == *c.checked {
			box = "X"
		}

		cursor := "  "
		if focused && i == c.internalIdx {
			cursor = "> "
		}

		sb.WriteString(fmt.Sprintf("\033[K%s[%s] %s", cursor, box, option))

		if i < len(c.options)-1 {
			sb.WriteString("\r\n")
		}
	}

	return sb.String()
}

func (c *Checkbox) HandleInput(key []byte) bool {
	if len(key) == 1 {
		switch {
		case key[0] == menu.KeySpace || key[0] == menu.KeyEnter:
			if c.checked != nil {
				*c.checked = c.internalIdx
				if c.onChange != nil {
					c.onChange(*c.checked)
				}
			}
			return true
		}
	}

	if len(key) == int(menu.KeySeqLen) && key[0] == menu.KeyEsc && key[1] == menu.KeyOpenBracket {
		switch key[2] {
		case menu.ArrowUp:
			if c.internalIdx > 0 {
				c.internalIdx--
				return true
			}
			return false
		case menu.ArrowDown:
			if c.internalIdx < len(c.options)-1 {
				c.internalIdx++
				return true
			}
			return false
		}
	}

	return false
}
