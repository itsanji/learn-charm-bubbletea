package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type ShopList struct {
	choices  []string         // item on the list
	cursor   int              // which item out cursor is pointing at
	selected map[int]struct{} // selected item
}

func initialModel() ShopList {
	return ShopList{
		choices:  []string{"Carrot", "Tofu", "Blueberries", "Orange"},
		cursor:   0,
		selected: make(map[int]struct{}),
	}
}

func (m ShopList) Init() tea.Cmd {
	return nil
}

func (m ShopList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor < len(m.choices) {
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

func (m ShopList) View() string {
	s := "What should we buy in the market? \n"
	for i, choice := range m.choices {
		// Is curosr pointing at this record ?
		pointer := " "
		if m.cursor == i {
			pointer = ">"
		}

		// Is this record selected
		isSelected := " "
		if _, ok := m.selected[i]; ok {
			isSelected = "X"
		}

		// Render the row
		row := fmt.Sprintf("%s [%s] %s \n", pointer, isSelected, choice)
		s += row
	}
	s += "\n\n\nPress Q to quit."
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("There's an error %v", err)
		os.Exit(1)
	}
}
