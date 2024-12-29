package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	btea "github.com/charmbracelet/bubbletea"
)

type SpinnerWaitModel struct {
	SpinnerWait spinner.Model
	quitting    bool
	err         error
}

func SpinnerWaitInitialModel() SpinnerWaitModel {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return SpinnerWaitModel{SpinnerWait: s}
}

func (m SpinnerWaitModel) Init() btea.Cmd {
	return m.SpinnerWait.Tick
}

func (m SpinnerWaitModel) Update(msg btea.Msg) (btea.Model, btea.Cmd) {
	switch msg := msg.(type) {
	case btea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, btea.Quit
		default:
			return m, nil
		}

	case error:
		m.err = msg
		return m, nil

	default:
		var cmd btea.Cmd
		m.SpinnerWait, cmd = m.SpinnerWait.Update(msg)
		return m, cmd
	}
}

func (m SpinnerWaitModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s Controll the password\n\n", m.SpinnerWait.View())
	if m.quitting {
		return str + "\n"
	}
	return str
}
