package analysis

import (
	"context"
	"os"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

type AnalysisRequest struct {
	Symbol   string
	Duration int
	Interval string
}

var (
	BINANCE_API_KEY    string
	BINANCE_SECRET_KEY string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
		os.Exit(1)
	}

	BINANCE_API_KEY = os.Getenv("BINANCE_API_KEY")
	BINANCE_SECRET_KEY = os.Getenv("BINANCE_SECRET_KEY")

	if BINANCE_API_KEY == "" || BINANCE_SECRET_KEY == "" {
		log.Error("Error getting Binance keys")
		os.Exit(1)
	}
}

func StartAnalysis(analysisRequest *AnalysisRequest) {
	log.Info("Starting Analysis...")
	log.Info("Data", "Symbol", analysisRequest.Symbol)
	log.Info("Data", "Duration", analysisRequest.Duration)
	log.Info("Data", "Interval", analysisRequest.Interval)

	client := binance.NewClient(BINANCE_API_KEY, BINANCE_SECRET_KEY)

	klines, err := fetchKlines(client, analysisRequest)
	if err != nil {
		log.Error("Error fetching Klines data", "err", err)
	}

	performAnalysis(klines)
}

func fetchKlines(client *binance.Client, analysisRequest *AnalysisRequest) ([]*binance.Kline, error) {
	log.Info("Fetching KLines Data...")
	now := time.Now()
	daysAgo := now.AddDate(0, 0, analysisRequest.Duration)

	klines, err := client.NewKlinesService().
		Symbol(analysisRequest.Symbol).
		Interval("1d").
		StartTime(daysAgo.UnixMilli()).
		EndTime(now.UnixMilli()).
		Do(context.Background())

	if err != nil {
		return nil, err
	}

	return klines, nil
}

func performAnalysis(klines []*binance.Kline) {
	_ = klines
	log.Info("Performing Analysis...")
}
