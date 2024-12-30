package main

import (
	"fmt"
	"os"
	"time"

	btea "github.com/charmbracelet/bubbletea"
	bf "github.com/mattemello/passwordCracker/bruteForce"
	tui "github.com/mattemello/passwordCracker/tui"
)

func main() {
	p := btea.NewProgram(tui.PasswordInitialModel())
	pass, err := p.Run()
	if err != nil {
		os.Exit(1)
	}

	if pass, ok := pass.(tui.PasswordModel); ok && pass.PasswordInput.Value() != "" {
		timePassed := make(chan time.Duration)
		go bf.BruteForcePass(pass.PasswordInput.Value(), timePassed)
		p = btea.NewProgram(tui.SpinnerWaitInitialModel(timePassed))

		_, err := p.Run()
		if err != nil {
			os.Exit(1)
		}

		fmt.Println(<-timePassed)
	}
}
