package main

import (
	"fmt"
	"os"
	"time"

	analysis "github.com/BrianMwangi21/anti-charts.git/pkg/analysis"
	cli "github.com/BrianMwangi21/anti-charts.git/pkg/cli"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
		os.Exit(1)
	}

	analysis.BINANCE_API_KEY = os.Getenv("BINANCE_API_KEY")
	analysis.BINANCE_SECRET_KEY = os.Getenv("BINANCE_SECRET_KEY")
	analysis.ALPACA_API_KEY = os.Getenv("ALPACA_API_KEY")
	analysis.ALPACA_SECRET_KEY = os.Getenv("ALPACA_SECRET_KEY")
	analysis.ALPACA_BASE_URL = os.Getenv("ALPACA_BASE_URL")
	perform_trades := os.Getenv("PERFORM_TRADES")
	special_cases := os.Getenv("SPECIAL_CASES")

	if analysis.BINANCE_API_KEY == "" || analysis.BINANCE_SECRET_KEY == "" {
		log.Error("Error getting Binance keys")
		os.Exit(1)
	}

	if analysis.ALPACA_API_KEY == "" || analysis.ALPACA_SECRET_KEY == "" || analysis.ALPACA_BASE_URL == "" {
		log.Error("Error getting Alpaca keys")
		os.Exit(1)
	}

	if special_cases == "True" {
		analysis.SPECIAL_CASES = true
	} else if special_cases == "False" {
		analysis.SPECIAL_CASES = false
	} else {
		log.Error("Error getting Special Cases key")
		os.Exit(1)
	}

	if perform_trades == "True" {
		analysis.PERFORM_TRADES = true
	} else if perform_trades == "False" {
		analysis.PERFORM_TRADES = false
	} else {
		log.Error("Error getting Perform Trades key")
		os.Exit(1)
	}
}

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
				seconds := currentTime.Second()

				if (minutes+1)%5 == 0 && seconds < 55 {
					break
				}

				time.Sleep(1 * time.Second)
			}
			log.Info(fmt.Sprintf("Time started %v", currentTime.Format(time.RFC3339)))

			if analysis.PERFORM_TRADES {
				ticker := time.NewTicker(10 * time.Second)
				go func() {
					for {
						select {
						case <-ticker.C:
							analysis.CheckMetrics()
						}
					}
				}()
			}

			analysis.StartAnalysis(analysisRequest)
		}
	}

}
