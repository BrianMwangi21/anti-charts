package analysis

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/charmbracelet/log"
	"github.com/cinar/indicator"
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
	ASSET              *indicator.Asset
	USER_INPUT_CHANNEL = make(chan string)
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

	saveData(klines)
	performAnalysis()
	performStrategies()
	RestartAnalysis()
}

func RestartAnalysis() {
	seconds := 300
	fmt.Println()
	fmt.Print("\033[s")

	go func() {
		var input string
		fmt.Scanln(&input)
		USER_INPUT_CHANNEL <- input
	}()

	for seconds > 0 {
		select {
		// Check if user pressed enter
		case userInput := <-USER_INPUT_CHANNEL:
			if userInput == "" {
				StartAnalysis(ANALYSIS_REQ)
				return
			}
		default:
			fmt.Printf("Restarting Analysis in %v seconds... (Press Enter at any point to skip wait)", seconds)
			time.Sleep(time.Second)
			seconds--
			fmt.Print("\033[u\033[K")
		}
	}

	StartAnalysis(ANALYSIS_REQ)
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

func saveData(klines []*binance.Kline) {
	OPENS := make([]float64, 0, len(klines))
	CLOSES := make([]float64, 0, len(klines))
	HIGHS := make([]float64, 0, len(klines))
	LOWS := make([]float64, 0, len(klines))
	VOLUMES := make([]float64, 0, len(klines))
	DATES := make([]time.Time, 0, len(klines))

	for _, k := range klines {
		if openPrice, err := strconv.ParseFloat(k.Open, 64); err == nil {
			OPENS = append(OPENS, openPrice)
		} else {
			log.Error("Error parsing open price: %v", err)
			return
		}

		if closePrice, err := strconv.ParseFloat(k.Close, 64); err == nil {
			CLOSES = append(CLOSES, closePrice)
		} else {
			log.Error("Error parsing close price: %v", err)
			return
		}

		if high, err := strconv.ParseFloat(k.High, 64); err == nil {
			HIGHS = append(HIGHS, high)
		} else {
			log.Error("Error parsing high price: %v", err)
			return
		}

		if low, err := strconv.ParseFloat(k.Low, 64); err == nil {
			LOWS = append(LOWS, low)
		} else {
			log.Error("Error parsing low price: %v", err)
			return
		}

		if volume, err := strconv.ParseFloat(k.Volume, 64); err == nil {
			VOLUMES = append(VOLUMES, volume)
		} else {
			log.Error("Error parsing volume: %v", err)
			return
		}

		closeTime := time.Unix(k.CloseTime/1000, 0)
		DATES = append(DATES, closeTime)
	}

	ASSET = &indicator.Asset{
		Date:    DATES,
		Opening: OPENS,
		Closing: CLOSES,
		High:    HIGHS,
		Low:     LOWS,
		Volume:  VOLUMES,
	}
}

func performAnalysis() {
	log.Info("Performing Indicator Analysis...")
	start := time.Now()
	defer trackTime("Indicator Analysis", start)
	var analysis []string

	analysis = append(analysis, performMACD(ASSET.Closing))
	analysis = append(analysis, performSMA(ASSET.Closing, ANALYSIS_REQ.Duration, LATEST_PRICE))
	analysis = append(analysis, performEMA(ASSET.Closing, ANALYSIS_REQ.Duration, LATEST_PRICE))
	analysis = append(analysis, performDEMA(ASSET.Closing, ANALYSIS_REQ.Duration, LATEST_PRICE))
	analysis = append(analysis, performTEMA(ASSET.Closing, ANALYSIS_REQ.Duration, LATEST_PRICE))
	analysis = append(analysis, performRSI(ASSET.Closing))
	analysis = append(analysis, performBB(ASSET.Closing, LATEST_PRICE))
	analysis = append(analysis, performMFI(ANALYSIS_REQ.Duration, ASSET.High, ASSET.Low, ASSET.Closing, ASSET.Volume))

	for _, value := range analysis {
		log.Info("INDICATORS", "result", value)
	}
}

func performStrategies() {
	log.Info("Performing Strategies Analysis...")
	start := time.Now()
	defer trackTime("Strategies Analysis", start)
	performAllStrategies(ASSET, ANALYSIS_REQ.Duration)
}

func ValidateInput(input []string) (*AnalysisRequest, error) {
	log.Debug("Validating input...")

	symbol := strings.ToUpper(input[0])
	if len(symbol) == 0 {
		return nil, errors.New("Symbol entry is invalid")
	}

	duration, err := strconv.Atoi(input[1])
	if err != nil {
		return nil, errors.New("Duration entry is invalid")
	}

	interval := input[2]
	pattern := `^(1m|3m|5m|15m|30m|1h|2h|4h|6h|8h|12h|1d|3d|1w|1M)$`
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, errors.New("Error compiling regex")
	}

	if !r.MatchString(interval) {
		return nil, errors.New("Interval entry is invalid")
	}

	return &AnalysisRequest{
		Symbol:   symbol,
		Duration: duration,
		Interval: interval,
	}, nil
}

func trackTime(analysisName string, start time.Time) {
	elapsed := time.Since(start)
	log.Info(fmt.Sprintf("Performing %v", analysisName), "time elapsed", elapsed)
}
