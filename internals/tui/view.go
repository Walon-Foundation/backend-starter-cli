package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// SelectionModel manages the state for the interactive selection
type SelectionModel struct {
	choices  []string      // available options
	cursor   int           // which option our cursor is pointing at
	selected map[int]bool  // which options are selected
	question string        // question to display above options
}

// InitialModel creates a new selection model with the given options and question
func InitialModel(question string, choices []string) SelectionModel {
	return SelectionModel{
		choices:  choices,
		selected: make(map[int]bool),
		question: question,
	}
}

func (m SelectionModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please"
	return nil
}

func (m SelectionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
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
			// Select the current choice
			m.selected[m.cursor] = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m SelectionModel) View() string {
	// The header with the question
	s := fmt.Sprintf("%s\n\n", m.question)

	// Iterate over our choices
	for i, choice := range m.choices {
		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if m.selected[i] {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit, ↑↓ to move, enter to select.\n"

	// Send the UI for rendering
	return s
}

// Helper function to run the selection and return the chosen option
func RunSelection(question string, options []string) ( string, error) {
	if len(options) == 0 {
		return "", fmt.Errorf("no options provided")
	}

	p := tea.NewProgram(InitialModel(question, options))
	
	// Run the program and get the final model
	finalModel, err := p.Run()
	if err != nil {
		return "", err
	}

	// Type assert back to our model
	if model, ok := finalModel.(SelectionModel); ok {
		// Find which option was selected
		for i := range model.choices {
			if model.selected[i] {
				return model.choices[i], nil
			}
		}
	}

	return "", fmt.Errorf("no option was selected")
}