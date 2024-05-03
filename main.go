package main

import (
	"os"

	"github.com/BrianMwangi21/anti-charts.git/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

func main() {
	if _, err := tea.NewProgram(utils.InitModel()).Run(); err != nil {
		log.Error("Could not start the program", "err", err)
		os.Exit(1)
	}
}
