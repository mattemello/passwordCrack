package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	btea "github.com/charmbracelet/bubbletea"
)

type PasswordModel struct {
	PasswordInput textinput.Model
	err           error
}

func PasswordInitialModel() PasswordModel {
	ti := textinput.New()
	ti.Placeholder = "Password to test"
	ti.EchoMode = textinput.EchoPassword
	ti.Focus()
	ti.Width = 20
	ti.CharLimit = 8

	return PasswordModel{
		PasswordInput: ti,
		err:           nil,
	}
}

func (m PasswordModel) Init() btea.Cmd {
	return textinput.Blink
}

func (m PasswordModel) Update(msg btea.Msg) (btea.Model, btea.Cmd) {
	var cmd btea.Cmd

	switch msg := msg.(type) {
	case btea.KeyMsg:
		switch msg.Type {
		case btea.KeyEsc:
			return m, btea.Quit
		case btea.KeyEnter:
			return m, btea.Quit
		}

	case error:
		m.err = msg
		return m, nil
	}

	m.PasswordInput, cmd = m.PasswordInput.Update(msg)
	return m, cmd
}

func (m PasswordModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	return fmt.Sprintf("insert the password you want to test: \n\n%s\n\n", m.PasswordInput.View())
}
