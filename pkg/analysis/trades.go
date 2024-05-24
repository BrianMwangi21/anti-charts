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
	SYMBOL := ANALYSIS_REQ.Base + DEFAULT_APLACA_QUOTE

	if action == indicator.BUY || action == indicator.SELL {
		client := getAlpacaClient()

		account, err := client.GetAccount()
		if err != nil {
			log.Error("Error getting account", "err", err)
			os.Exit(1)
		}
		log.Info("TRADING", "accountBuyingPower", account.BuyingPower)
		log.Info("TRADING", "accountBalanceChange", account.Equity.Sub(account.LastEquity))
		log.Info("TRADING", "symbol", SYMBOL)

		if action == indicator.BUY {
			performBuyTrade(client, SYMBOL)
		} else if action == indicator.SELL {
			performSellTrade(client, SYMBOL)
		}
	} else {
		log.Info("TRADING", "action", "HOLD. Doing nothing for now")
	}
}

func performBuyTrade(client *alpaca.Client, symbol string) {
	log.Info("Performing Buy Trade...")
	start := time.Now()
	defer trackTime("Performing Buy Trade", start)

	NOTIONAL := decimal.NewFromInt(100)
	order, err := client.PlaceOrder(alpaca.PlaceOrderRequest{
		Symbol:      symbol,
		Notional:    &NOTIONAL,
		Side:        alpaca.Buy,
		Type:        alpaca.Market,
		TimeInForce: alpaca.IOC,
	})

	if err != nil {
		log.Error("Error placing buy order", "err", err)
	} else {
		log.Info("TRADING", "buyOrderPlaced", order.ID)
	}
}

func performSellTrade(client *alpaca.Client, symbol string) {
	log.Info("Performing Sell Trade...")
	start := time.Now()
	defer trackTime("Performing Sell Trade", start)

	position, err := client.GetPosition(symbol)
	if err != nil {
		log.Error("Error getting position", "err", err)
	} else {
		QTY := position.QtyAvailable
		order, err := client.PlaceOrder(alpaca.PlaceOrderRequest{
			Symbol:      symbol,
			Qty:         &QTY,
			Side:        alpaca.Sell,
			Type:        alpaca.Market,
			TimeInForce: alpaca.IOC,
		})

		if err != nil {
			log.Error("Error placing sell order", "err", err)
		} else {
			log.Info("TRADING", "sellOrderPlaced", order.ID)
		}
	}
}

func performCleanup() {
	symbol := ANALYSIS_REQ.Base + DEFAULT_APLACA_QUOTE
	client := getAlpacaClient()

	log.Info("Performing Cleanup...")
	start := time.Now()
	defer trackTime("Performing Cleanup", start)
	performSellTrade(client, symbol)
}
