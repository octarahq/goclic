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
		return fmt.Errorf("impossible de lancer un menu sans composants")
	}

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return fmt.Errorf("erreur initialisation mode brut: %w", err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	m.running = true
	firstRender := true

	for m.running {
		if !firstRender {
			for i := 0; i < len(m.components); i++ {
				fmt.Print("\033[F")
			}
		}
		firstRender = false

		for i, comp := range m.components {
			isFocused := (i == m.idX)
			fmt.Print("\033[K" + comp.Render(isFocused) + "\r\n")
		}

		buf := make([]byte, 3)
		n, err := os.Stdin.Read(buf)
		if err != nil {
			return err
		}

		if n == 1 && buf[0] == 3 {
			break
		}

		if n == 3 && buf[0] == 27 && buf[1] == 91 {
			switch buf[2] {
			case 65:
				if m.idX > 0 {
					m.idX--
				} else {
					m.idX = len(m.components) - 1
				}
				continue

			case 66:
				if m.idX < len(m.components)-1 {
					m.idX++
				} else {
					m.idX = 0
				}
				continue
			}
		}

		activeComp := m.components[m.idX]
		if interactive, ok := activeComp.(InteractiveComponent); ok {
			interactive.HandleInput(buf[:n])
		}
	}

	return nil
}
