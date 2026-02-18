package prompt

import (
    "fmt"
    "os"

    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type ProjectData struct {
    Name string
}

type model struct {
    textInput textinput.Model
    err       error
}

func initialModel() model {
    ti := textinput.New()
    ti.Placeholder = "My Awesome Project"
    ti.Focus()
    ti.CharLimit = 156
    ti.Width = 20

    return model{
        textInput: ti,
        err:       nil,
    }
}

func (m model) Init() tea.Cmd {
    return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
            return m, tea.Quit
        }
    case error:
        m.err = msg
        return m, nil
    }

    m.textInput, cmd = m.textInput.Update(msg)
    return m, cmd
}

func (m model) View() string {
    return fmt.Sprintf(
        "What is the name of your project?\n\n%s\n\n%s",
        m.textInput.View(),
        "(esc to quit)",
    ) + "\n"
}

// Run starts the interactive UI and returns the collected data.
func Run() (*ProjectData, error) {
    p := tea.NewProgram(initialModel())
    m, err := p.Run()
    if err != nil {
        return nil, err
    }

    if finalModel, ok := m.(model); ok {
        if finalModel.textInput.Value() == "" {
             return nil, fmt.Errorf("no project name provided")
        }
        return &ProjectData{Name: finalModel.textInput.Value()}, nil
    }

    return nil, fmt.Errorf("could not retrieve model")
}
