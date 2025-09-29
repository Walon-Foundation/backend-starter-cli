package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
)

// InputModel manages the state for the text input
type InputModel struct {
	textInput textinput.Model
	err       error
	question  string
}

// InitialInputModel creates a new input model with the given question
func InitialInputModel(question string) InputModel {
	ti := textinput.New()
	ti.Placeholder = "Type your answer..."
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 50

	return InputModel{
		textInput: ti,
		err:       nil,
		question:  question,
	}
}

func (m InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			m.textInput.SetValue("") // Return empty string if user cancels
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m InputModel) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		m.question,
		m.textInput.View(),
		"(press enter to confirm)",
	) + "\n"
}

// RunInput displays a text input prompt and returns the user's input
func RunInput(question string) (string, error) {
	p := tea.NewProgram(InitialInputModel(question))
	
	// Run the program and get the final model
	finalModel, err := p.Run()
	if err != nil {
		return "", err
	}

	// Type assert back to our model and get the input value
	if model, ok := finalModel.(InputModel); ok {
		return model.textInput.Value(), nil
	}

	return "", fmt.Errorf("failed to get input")
}