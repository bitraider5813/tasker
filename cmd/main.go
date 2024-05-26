package main


import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)


type model struct {
	choices []string
	cursor  int
	selected map[int]struct{} 
}


// This is a private function that will return a model
func initialModel() model {
	return model{
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		selected: make(map[int]struct{}),
	}
}


func (m model) Init() tea.Cmd {
	// start by returning nil which is the same as "no command"
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
        // Return the updated model to the Bubble Tea runtime for processing.
	// Note that we are not returning a command
	return m, nil
}


func (m model) View() string {
	// The header
	s := "Tasker To-Do List: My First To-Do List\n\n"

	for i, choice := range m.choices {
		// is the cursor pointing at this choice?
		cursor := " "  // no cursor
		if m.cursor == i {
			cursor = ">" // set the cursor!
		}

		// is this choice selcted?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" //selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// the footer
	s += "\nPress q to quit.\n"

	// send the UI for rendering
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alias, there's been an error: %v", err)
		os.Exit(1)
	}
}
