package components

import (
	"fmt"

	menu "github.com/octarahq/goclic"
)

type Button struct {
	label   string
	onClick func()
}

func NewButton(label string, onClick func()) *Button {
	return &Button{
		label:   label,
		onClick: onClick,
	}
}

func (d *Button) GetName() string {
	return "button"
}

func (d *Button) Height() int {
	return 1
}

func (d *Button) IsSelectable() bool {
	return true
}

func (l *Button) Render(focused bool) string {
	if focused {
		return fmt.Sprintf("> [%s]", l.label)
	}
	return fmt.Sprintf("  [%s]", l.label)
}

func (l *Button) HandleInput(key []byte) bool {
	if len(key) == 1 && (key[0] == menu.KeyEnter || key[0] == menu.KeySpace) {
		l.onClick()
		return true
	}
	return false
}
