package menu

type Component interface {
	GetName() string
	Render(focused bool) string
	IsSelectable() bool
	Height() int
}

type InteractiveComponent interface {
	Component
	HandleInput(key []byte) bool
}
