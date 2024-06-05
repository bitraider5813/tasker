package main


import (
	"fmt"
	"os"
	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
)


type model struct {
	tasks []string
	cursor  int
	completed map[int]struct{} 
}


// This is a private function that will return a model
func initialModel() model {
	return model{
		tasks: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		completed: make(map[int]struct{}),
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
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.completed[m.cursor]
			if ok {
				delete(m.completed, m.cursor)
			} else {
				m.completed[m.cursor] = struct{}{}
			}
		}
	}
        // Return the updated model to the Bubble Tea runtime for processing.
	// Note that we are not returning a command
	return m, nil
}


func (m model) View() string {
	// The header
	var style = lipgloss.NewStyle().
	  Bold(true).
	  Foreground(lipgloss.Color("#3C3C3C")).
	  Background(lipgloss.Color("86")).
	  PaddingTop(0).
	  PaddingBottom(0).
	  PaddingLeft(4).
	  Width(52)
	
	
	var selectedStyle = lipgloss.NewStyle().
	  Foreground(lipgloss.Color("#3C3C3C")).
	  PaddingTop(0).
	  PaddingBottom(0).
	  Margin(1)

	var normalStyle = lipgloss.NewStyle().
	  Foreground(lipgloss.Color("#04B575")).
	  PaddingTop(0).
	  PaddingBottom(0).
	  Margin(1)


	s := fmt.Sprintf(style.Render("Tasker To-Do List: My First To-Do List"))
	s += fmt.Sprintf("%s%s", "\n", "\n")

	for i, task := range m.tasks {
		// is the cursor pointing at this choice?
		cursor := " "  // no cursor
		if m.cursor == i {
			cursor = ">" // set the cursor!
		}

		// is this choice selcted?
		checked := " " // not selected
		if _, ok := m.completed[i]; ok {
			checked = "\u2713" //selected!
		}

		// Render the row
		if checked == "\u2713" {
			checkedString := fmt.Sprintf("%s [%s] %s\n", cursor, checked, task)
			s += fmt.Sprintf(selectedStyle.Render(checkedString))
		} else {
			normalString := fmt.Sprintf("%s [%s] %s\n", cursor, checked, task)
			s += fmt.Sprintf(normalStyle.Render(normalString))
		}
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
