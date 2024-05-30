package main

import (
	"os"
	"time"

	analysis "github.com/BrianMwangi21/anti-charts.git/pkg/analysis"
	cli "github.com/BrianMwangi21/anti-charts.git/pkg/cli"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

func main() {
	p := tea.NewProgram(cli.InitModel())

	if model, err := p.Run(); err != nil {
		log.Error("Error starting the program", "err", err)
		os.Exit(1)
	} else {
		initModel := model.(cli.Model)

		if initModel.Submitted {
			analysisRequest, err := analysis.ValidateInput(initModel.Values)
			if err != nil {
				log.Error("Error validating input", "err", err)
				os.Exit(1)
			}

			ticker := time.NewTicker(time.Duration(analysis.CHECK_METRICS_WAIT_SECONDS) * time.Second)
			go func() {
				for {
					select {
					case <-ticker.C:
						analysis.CheckMetrics()
					}
				}
			}()

			analysis.StartAnalysis(analysisRequest)

		}
	}

}
