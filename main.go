package main

import (
	"fmt"
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

			// Start execution at multiples of 5 for symmetry
			var currentTime time.Time
			log.Info("Checking time...")
			for {
				currentTime = time.Now()
				minutes := currentTime.Minute()

				if minutes%5 == 0 {
					break
				}

				time.Sleep(10 * time.Second)
			}
			log.Info(fmt.Sprintf("Time started %v", currentTime.Format(time.RFC3339)))

			ticker := time.NewTicker(10 * time.Second)
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
