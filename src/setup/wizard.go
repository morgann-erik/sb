package setup

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/morgann-erik/sb/core"
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

	inputs[userName] = textinput.New()
	inputs[userName].Placeholder = "username"

	m := model{inputs: inputs, focusedField: 0}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

type model struct {
	focusedField int
	inputs       []textinput.Model
}

func (m *model) prevInput() {
	m.focusedField--
	if m.focusedField < 0 {
		m.focusedField = len(m.inputs) - 1
	}
}

func (m *model) nextInput() {
	m.focusedField = (m.focusedField + 1) % len(m.inputs)
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlQ:
		case tea.KeyCtrlC:
		case tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyShiftTab:
			m.prevInput()
		case tea.KeyTab:
			m.nextInput()
		case tea.KeyEnter:
            createConfigFile(m)
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
		m.inputs[i].Blur()
	}
	m.inputs[m.focusedField].Focus()

	return m, tea.Batch(cmds...)
}

func (m model) View() string {

	return fmt.Sprintf(` 
    Silverback: setup wizard
    === == = == === == = == ===
    [ctrl+q] exit | [tab] next | [shift+tab] prev | [enter] continue 
    --- -- - -- --- -- - -- ---

    %s
    %s`, m.inputs[host].View(), m.inputs[userName].View())
}

func createConfigFile(m model) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	//check if .config exists if not create
    configDir := path.Join(homeDir, core.ConfigDir)
    _, err = os.Stat(configDir)
    if os.IsNotExist(err) {
        os.Mkdir(configDir, os.ModePerm)
    }

	// check if sb exists
    sbDir := path.Join(configDir, core.SbDir)
    _, err = os.Stat(sbDir)
    if os.IsNotExist(err) {
        os.Mkdir(sbDir, os.ModePerm)
    }

	// create config file
    cfgPath := path.Join(sbDir, core.ConfigFile)
    if err = os.WriteFile(cfgPath, []byte(""), 0666); err != nil {
        log.Fatal(err)
        return err
    }

	// create templates

	// create projects

    return nil
}
