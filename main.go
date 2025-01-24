package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string         // items on the display list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func initialModel() model {
	return model{
		choices:  []string{"کلمبیا", "کنیا", "بلند ۸۰ درصد"},
		cursor:   0,
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

		}
	}
	return m, nil
}

func (m model) View() string {
	s := "قهوه خود را انتخاب کنید\n\n"

	for i, choice := range m.choices {
		cursor := " "

		// Line pointer
		if m.cursor == i {
			cursor = ">"
		}

		// selected
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "X"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)

	}

	s += "\nPress q to quit.\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	_, err := p.Run()
	if err != nil {
		for i := 0; i < 200; i++ {
			fmt.Print("ERROR \t")
		}
		os.Exit(1)
	}
}
