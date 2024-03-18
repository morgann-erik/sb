package setup

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func startWizard() {
    m := model{
        userName: textinput.New(),
    }
    m.userName.Placeholder = "username"
    m.userName.Focus()

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

type model struct {
    userName textinput.Model
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

    var uMsg tea.Cmd
    m.userName, uMsg = m.userName.Update(msg)

	return m, uMsg
}

func (m model) View() string {

	return fmt.Sprintf("\n\n Hello world!\n\n%s", m.userName.View())
}
