package analysis

import (
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/cinar/indicator"
)

type AnalysisRequest struct {
	Base     string
	Duration int
	Interval string
}

var (
	DEFAULT_QUOTE            = "USDT"
	DEFAULT_APLACA_QUOTE     = "USD"
	WAIT_SECONDS             = 300
	DEFAULT_NOTIONAL_VALUE   = int64(250)
	DEFAULT_PORTFOLIO_CHANGE = int64(3)
	BINANCE_API_KEY          string
	BINANCE_SECRET_KEY       string
	ALPACA_API_KEY           string
	ALPACA_SECRET_KEY        string
	ALPACA_BASE_URL          string
	ANALYSIS_REQ             *AnalysisRequest
	LATEST_PRICE             float64
	ASSET                    *indicator.Asset
	USER_INPUT_CHANNEL       = make(chan string)
	ALPACA_CLIENT            *alpaca.Client
	ONCE_ALPACA              sync.Once
	BINANCE_CLIENT           *binance.Client
	ONCE_BINANCE             sync.Once
	LAST_ACTIONS             []indicator.Action
	DUMP_STOCK               bool
	HOLD_STOCK               bool
)

func getBinanceClient() *binance.Client {
	ONCE_BINANCE.Do(func() {
		BINANCE_CLIENT = binance.NewClient(BINANCE_API_KEY, BINANCE_SECRET_KEY)
	})
	return BINANCE_CLIENT
}

func getAlpacaClient() *alpaca.Client {
	ONCE_ALPACA.Do(func() {
		ALPACA_CLIENT = alpaca.NewClient(alpaca.ClientOpts{
			APIKey:     ALPACA_API_KEY,
			APISecret:  ALPACA_SECRET_KEY,
			BaseURL:    ALPACA_BASE_URL,
			RetryDelay: 3 * time.Second,
		})
	})
	return ALPACA_CLIENT
}
