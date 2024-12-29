package main

import (
	"os"

	btea "github.com/charmbracelet/bubbletea"
	tui "github.com/mattemello/passwordCracker/tui"
)

func main() {
	p := btea.NewProgram(tui.PasswordInitialModel())
	pass, err := p.Run()
	if err != nil {
		os.Exit(1)
	}

	if pass, ok := pass.(tui.PasswordModel); ok && pass.PasswordInput.Value() != "" {
		go testPassword(pass.PasswordInput.Value())
		p = btea.NewProgram(tui.SpinnerWaitInitialModel())
		_, err := p.Run()
		if err != nil {
			os.Exit(1)
		}
	}
}
