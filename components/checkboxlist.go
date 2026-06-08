package components

import (
	"fmt"
	"strings"

	menu "github.com/octarahq/goclic"
)

type CheckboxList struct {
	label       string
	options     []string
	checked     *[]bool
	internalIdx int
	onChange    func(idx int)
}

type CheckboxListOption func(*CheckboxList)

func WithCheckboxListOnChange(onChange func(idx int)) CheckboxListOption {
	return func(c *CheckboxList) {
		c.onChange = onChange
	}
}

func NewCheckboxList(label string, options []string, checked *[]bool, opts ...CheckboxListOption) *CheckboxList {
	if len(*checked) == 0 {
		*checked = make([]bool, len(options))
	}
	c := &CheckboxList{
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

func (c *CheckboxList) GetName() string    { return "checkboxlist" }
func (c *CheckboxList) IsSelectable() bool { return true }

func (c *CheckboxList) Height() int {
	return 1 + len(c.options)
}

func (c *CheckboxList) Render(focused bool) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("  %s :\r\n", c.label))

	for i, option := range c.options {
		box := " "
		if (*c.checked)[i] {
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

func (c *CheckboxList) HandleInput(key []byte) bool {
	if len(key) == 1 {
		switch {
		case key[0] == menu.KeySpace || key[0] == menu.KeyEnter:
			(*c.checked)[c.internalIdx] = !(*c.checked)[c.internalIdx]
			if c.onChange != nil {
				c.onChange(c.internalIdx)
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
