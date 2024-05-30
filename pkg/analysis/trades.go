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

		if action == indicator.BUY {
			performBuyTrade(client, account.BuyingPower, SYMBOL)
		} else if action == indicator.SELL {
			performSellTrade(client, SYMBOL)
		}
	} else {
		log.Info("TRADING", "action", "HOLD. Doing nothing for now")
	}
}

func performBuyTrade(client *alpaca.Client, buyingPower decimal.Decimal, symbol string) {
	log.Info("Performing Buy Trade...")
	start := time.Now()
	defer trackTime("Performing Buy Trade", start)

	notional := decimal.NewFromInt(DEFAULT_NOTIONAL_VALUE)

	if buyingPower.GreaterThan(notional) {
		order, err := client.PlaceOrder(alpaca.PlaceOrderRequest{
			Symbol:      symbol,
			Notional:    &notional,
			Side:        alpaca.Buy,
			Type:        alpaca.Market,
			TimeInForce: alpaca.IOC,
		})

		if err != nil {
			log.Error("Error placing buy order", "err", err)
		} else {
			log.Info("TRADING", "buyOrderPlaced", order.ID)
		}
	} else {
		log.Info("TRADING", "action", "NO BUY. Not enough buying power")
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
		qty := position.QtyAvailable
		marketValue := position.MarketValue
		notional := decimal.NewFromInt(DEFAULT_NOTIONAL_VALUE)

		if marketValue.GreaterThan(notional) {
			log.Info("TRADING", "placingSellUsing", "Notional")
			order, err := client.PlaceOrder(alpaca.PlaceOrderRequest{
				Symbol:      symbol,
				Notional:    &notional,
				Side:        alpaca.Sell,
				Type:        alpaca.Market,
				TimeInForce: alpaca.IOC,
			})

			if err != nil {
				log.Error("Error placing sell order using NOTIONAL", "err", err)
			} else {
				log.Info("TRADING", "sellOrderPlaced", order.ID)
			}
		} else {
			log.Info("TRADING", "placingSellUsing", "QtyAvailable")
			order, err := client.PlaceOrder(alpaca.PlaceOrderRequest{
				Symbol:      symbol,
				Qty:         &qty,
				Side:        alpaca.Sell,
				Type:        alpaca.Market,
				TimeInForce: alpaca.IOC,
			})

			if err != nil {
				log.Error("Error placing sell order using QTY", "err", err)
			} else {
				log.Info("TRADING", "sellOrderPlaced", order.ID)
			}
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

func CheckMetrics() {
	client := getAlpacaClient()
	account, err := client.GetAccount()
	if err != nil {
		log.Error("Error getting account", "err", err)
		os.Exit(1)
	}

	accountBalanceChange := account.Equity.Sub(account.LastEquity)
	accountPercentageChange := accountBalanceChange.Div(account.LastEquity).Mul(decimal.NewFromInt(100))

	log.Info("================")
	log.Info("CHECKING METRICS", "accountBuyingPower", account.BuyingPower)
	log.Info("CHECKING METRICS", "accountPortfolioValue", account.PortfolioValue)
	log.Info("CHECKING METRICS", "accountBalanceChange", accountBalanceChange)
	log.Info("CHECKING METRICS", "accountPercentageChange", accountPercentageChange)

	if accountPercentageChange.GreaterThan(decimal.NewFromInt(DEFAULT_PORTFOLIO_CHANGE)) {
		log.Info("CHECKING METRICS", "We made some good money. Exiting now.")
		os.Exit(1)
	}
	log.Info("================")
}
