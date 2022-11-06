package main

import (
	"fmt"
	"os"

	slicemethod "todo_list/sliceMethod"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type Item struct {
	id    uuid.UUID
	title string
}

type ShopList struct {
	choices  []Item      // item on the list
	cursor   int         // which item out cursor is pointing at
	selected []uuid.UUID // selected item

	// appState string // View | Add
	// textInput textinput.Model // Text Input
}

func initialModel() ShopList {
	return ShopList{
		choices: []Item{
			{
				id:    uuid.New(),
				title: "Carrot",
			},
			{
				id:    uuid.New(),
				title: "BlueBerry",
			},
			{
				id:    uuid.New(),
				title: "Oranges",
			},
			{
				id:    uuid.New(),
				title: "Lemon",
			},
		},
		cursor:   0,
		selected: make([]uuid.UUID, 0),
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
			compareFunction := func(item uuid.UUID) bool {
				return item != m.choices[m.cursor].id
			}
			// Remove current choice from selected list
			if inList := slicemethod.IsInList(m.selected, m.choices[m.cursor].id); inList {
				m.selected = slicemethod.SliceFilter(m.selected, compareFunction)
			} else {
				// Add current choice to selected list
				m.selected = append(m.selected, m.choices[m.cursor].id)
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
		if ok := slicemethod.IsInList(m.selected, choice.id); ok {
			isSelected = "X"
		}

		// Render the row
		row := fmt.Sprintf("%s [%s] %s \n", pointer, isSelected, choice.title)
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
