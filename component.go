package menu

type Component interface {
	GetName() string
	Render(focused bool) string
}

type InteractiveComponent interface {
	Component
	HandleInput(key []byte) bool
}
