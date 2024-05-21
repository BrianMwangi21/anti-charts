package analysis

import "github.com/cinar/indicator"

type AnalysisRequest struct {
	Base     string
	Duration int
	Interval string
}

var (
	DEFAULT_QUOTE      = "USDT"
	BINANCE_API_KEY    string
	BINANCE_SECRET_KEY string
	ALPACA_API_KEY     string
	ALPACA_SECRET_KEY  string
	ALPACA_BASE_URL    string
	ANALYSIS_REQ       *AnalysisRequest
	LATEST_PRICE       float64
	ASSET              *indicator.Asset
	USER_INPUT_CHANNEL = make(chan string)
)
