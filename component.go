package menu

type Component interface {
	GetName() string
	Render(focused bool) string
	IsSelectable() bool
}

type InteractiveComponent interface {
	Component
	HandleInput(key []byte) bool
}
