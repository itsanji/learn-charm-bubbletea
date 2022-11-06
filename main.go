package main

import (
	"fmt"
	"os"

	slicemethod "todo_list/sliceMethod"

	"github.com/charmbracelet/bubbles/textinput"
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

	appState  string          // View | Add
	textInput textinput.Model // Text Input
}

func (m ShopList) Init() tea.Cmd {
	return nil
}

func (m ShopList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c": // Quit
			return m, tea.Quit
		case "q":
			if m.appState == "view" {
				return m, tea.Quit
			}
		case "up", "k": // Goes Up
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j": // Goes Down
			if m.cursor < len(m.choices) {
				m.cursor++
			}
		case "a", "A": // Change to Add Mode
			if m.appState == "view" {
				m.appState = "add"
				return m, textinput.Blink
			}
		case "d", "D": // Remove Item from list
			if m.appState == "view" {
				m.choices = slicemethod.SliceFilter(m.choices, func(choice Item) bool {
					return choice.id != m.choices[m.cursor].id
				})
			}
		case "enter": // Check
			if m.appState == "view" {
				// Toggle Todo Status
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
			} else {
				m.choices = append(m.choices, Item{
					id:    uuid.New(),
					title: m.textInput.Value(),
				})
				m.appState = "view"
				m.textInput.Reset()
			}
		}
	}
	if m.appState == "add" {
		m.textInput, cmd = m.textInput.Update(msg)
	}
	return m, cmd
}

func (m ShopList) View() string {
	s := ""
	if m.appState == "add" {
		s = fmt.Sprintf("New Todo:\n %s", m.textInput.View())
	} else {
		s = "Here is your Todo List: \n"
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
		s += "\n\nPress A to Add New Todo. \nPress D To Delete Current Todo \n\n\nPress Q to quit."
	}
	return s
}

func initialModel() ShopList {
	ti := textinput.New()
	ti.Placeholder = "New Todo.."
	ti.Focus()

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
		cursor:    0,
		appState:  "view",
		textInput: ti,
		selected:  make([]uuid.UUID, 0),
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("There's an error %v", err)
		os.Exit(1)
	}
}
