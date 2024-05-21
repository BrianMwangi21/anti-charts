package analysis

import (
	"os"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/charmbracelet/log"
	"github.com/cinar/indicator"
	"github.com/shopspring/decimal"
)

func performTrade(action indicator.Action) {
	log.Info("Performing Trade...")
	start := time.Now()
	defer trackTime("Performing Trade", start)
	SYMBOL := ANALYSIS_REQ.Base + "/USD"

	if action == indicator.BUY || action == indicator.SELL {
		var SIDE alpaca.Side

		if action == indicator.BUY {
			SIDE = alpaca.Buy
		} else if action == indicator.SELL {
			SIDE = alpaca.Sell
		}

		client := alpaca.NewClient(alpaca.ClientOpts{
			APIKey:    ALPACA_API_KEY,
			APISecret: ALPACA_SECRET_KEY,
			BaseURL:   ALPACA_BASE_URL,
		})

		account, err := client.GetAccount()
		if err != nil {
			log.Error("Error getting account", "err", err)
			os.Exit(1)
		}
		log.Info("TRADING", "accountBuyingPower", account.BuyingPower)
		log.Info("TRADING", "accountBalanceChange", account.Equity.Sub(account.LastEquity))
		log.Info("TRADING", "symbol", SYMBOL)

		NOTIONAL := decimal.NewFromInt(100)
		order, err := client.PlaceOrder(alpaca.PlaceOrderRequest{
			Symbol:      SYMBOL,
			Notional:    &NOTIONAL,
			Side:        SIDE,
			Type:        alpaca.Market,
			TimeInForce: alpaca.IOC,
		})

		if err != nil {
			log.Error("Error placing order", "err", err)
		}

		log.Info("TRADING", "orderPlaced", order.ID)
	} else {
		log.Info("TRADING", "action", "HOLD. Doing nothing for now")
	}
}
