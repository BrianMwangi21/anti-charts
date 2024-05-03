package analysis

import (
	"context"
	"os"
	"strconv"
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
	ANALYSIS_REQ       *AnalysisRequest
	LATEST_PRICE       float64
	CLOSES             []float64
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
	ANALYSIS_REQ = analysisRequest
	log.Info("Starting Analysis...")

	client := binance.NewClient(BINANCE_API_KEY, BINANCE_SECRET_KEY)
	fetchLatestPrice(client, ANALYSIS_REQ.Symbol)

	log.Info("Data", "Symbol", ANALYSIS_REQ.Symbol)
	log.Info("Data", "Duration", ANALYSIS_REQ.Duration)
	log.Info("Data", "Interval", ANALYSIS_REQ.Interval)
	log.Info("Data", "Latest Price", LATEST_PRICE)

	klines, err := fetchKlines(client, analysisRequest)
	if err != nil {
		log.Error("Error fetching Klines data", "err", err)
		os.Exit(1)
	}

	saveCloses(klines)
	performAnalysis()
}

func fetchLatestPrice(client *binance.Client, symbol string) {
	log.Info("Fetching Latest Price...")

	symbol_price, err := client.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		log.Error("Error fetching latest symbol price", "err", err)
		os.Exit(1)
	}

	LATEST_PRICE, err = strconv.ParseFloat(symbol_price[0].Price, 64)
	if err != nil {
		log.Error("Error converting latest symbol price", "err", err)
		os.Exit(1)
	}
}

func fetchKlines(client *binance.Client, analysisRequest *AnalysisRequest) ([]*binance.Kline, error) {
	log.Info("Fetching KLines Data...")
	now := time.Now()
	daysAgo := now.AddDate(0, 0, (analysisRequest.Duration * -1))

	klines, err := client.NewKlinesService().
		Symbol(analysisRequest.Symbol).
		Interval(analysisRequest.Interval).
		StartTime(daysAgo.UnixMilli()).
		EndTime(now.UnixMilli()).
		Do(context.Background())

	if err != nil {
		return nil, err
	}

	return klines, nil
}

func saveCloses(klines []*binance.Kline) {
	CLOSES = make([]float64, 0, len(klines))
	for _, k := range klines {
		if closePrice, err := strconv.ParseFloat(k.Close, 64); err == nil {
			CLOSES = append(CLOSES, closePrice)
		} else {
			log.Error("Error parsing close price: %v", err)
			return
		}
	}
}

func performAnalysis() {
	log.Info("Performing Analysis...")
	var analysis []string

	analysis = append(analysis, performMACD(CLOSES))
	analysis = append(analysis, performSMA(CLOSES, ANALYSIS_REQ.Duration, LATEST_PRICE))
	analysis = append(analysis, performEMA(CLOSES, ANALYSIS_REQ.Duration, LATEST_PRICE))
	analysis = append(analysis, performDEMA(CLOSES, ANALYSIS_REQ.Duration, LATEST_PRICE))
	analysis = append(analysis, performTEMA(CLOSES, ANALYSIS_REQ.Duration, LATEST_PRICE))

	log.Info("Preparing Analysis Results...")
	for _, value := range analysis {
		log.Info("Result", "data", value)
	}
}
