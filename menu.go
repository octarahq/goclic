package menu

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

type Menu struct {
	components []Component
	idX        int
	running    bool
}

func NewMenu() *Menu {
	return &Menu{
		components: make([]Component, 0),
		idX:        0,
		running:    false,
	}
}

func (m *Menu) Add(c Component) {
	m.components = append(m.components, c)
}

func (m *Menu) Stop() {
	m.running = false
}

func (m *Menu) Start() error {
	if len(m.components) == 0 {
		return fmt.Errorf("Cannot start an empty menu")
	}

	m.idX = -1
	for i, comp := range m.components {
		if comp.IsSelectable() {
			m.idX = i
			break
		}
	}

	if m.idX == -1 {
		m.idX = 0
	}

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return fmt.Errorf("raw mode initialization error: %w", err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	m.running = true
	firstRender := true

	for m.running {
		if !firstRender {
			linesToMove := 0
			for _, comp := range m.components {
				if multi, ok := comp.(interface{ Height() int }); ok {
					linesToMove += multi.Height()
				} else {
					linesToMove += 1
				}
			}

			for i := 0; i < linesToMove; i++ {
				fmt.Print("\033[F")
			}
		}
		firstRender = false

		for i, comp := range m.components {
			isFocused := (i == m.idX)
			fmt.Print("\033[K" + comp.Render(isFocused) + "\r\n")
		}

		buf := make([]byte, 1024)
		n, err := os.Stdin.Read(buf)
		if err != nil {
			return err
		}

		if n == 1 && buf[0] == KeyCtrlC {
			break
		}

		activeComp := m.components[m.idX]
		if interactive, ok := activeComp.(InteractiveComponent); ok {
			if interactive.HandleInput(buf[:n]) {
				continue
			}
		}

		if n == 3 && buf[0] == 27 && buf[1] == 91 {
			switch buf[2] {
			case ArrowUp:
				for {
					if m.idX > 0 {
						m.idX--
					} else {
						m.idX = len(m.components) - 1
					}
					if m.components[m.idX].IsSelectable() {
						break
					}
				}
				continue

			case ArrowDown:
				for {
					if m.idX < len(m.components)-1 {
						m.idX++
					} else {
						m.idX = 0
					}
					if m.components[m.idX].IsSelectable() {
						break
					}
				}
				continue
			}
		}
	}

	return nil
}
