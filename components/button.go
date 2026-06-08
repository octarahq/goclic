package components

import (
	"fmt"

	menu "github.com/octarahq/goclic"
)

type Button struct {
	label   string
	onClick func()
}

type ButtonOption func(*Button)

func WithButtonOnClick(onClick func()) ButtonOption {
	return func(b *Button) {
		b.onClick = onClick
	}
}

func NewButton(label string, opts ...ButtonOption) *Button {
	b := &Button{
		label: label,
	}
	for _, opt := range opts {
		opt(b)
	}
	return b
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
