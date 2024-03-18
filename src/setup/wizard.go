package setup

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	host = iota
	userName
)

func startWizard() {
	var inputs []textinput.Model = make([]textinput.Model, 2)
    inputs[host] = textinput.New()
    inputs[host].Placeholder = "host"
    inputs[host].Focus()

    inputs[userName].Placeholder = "username"

	m := model{inputs: inputs}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

type model struct {
	inputs []textinput.Model
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {

	return fmt.Sprintf(` 
    Silverback: setup wizard\n
    === == = == === == = == ===

    %s
    %s`, m.inputs[host].View(), m.inputs[userName].View())
}
