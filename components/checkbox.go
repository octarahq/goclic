package components

import (
	"fmt"
	"strings"

	menu "github.com/octarahq/goclic"
)

type Checkbox struct {
	label       string
	options     []string
	checked     int
	internalIdx int
}

func NewCheckbox(label string, options []string, checked *int) *Checkbox {
	return &Checkbox{
		label:       label,
		options:     options,
		checked:     *checked,
		internalIdx: 0,
	}
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
		if i == c.checked {
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
			c.checked = c.internalIdx
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
