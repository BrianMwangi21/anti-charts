package analysis

import (
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	"github.com/cinar/indicator"
)

const TREND_PERIOD = 5

func performAllStrategies(asset *indicator.Asset, period int) {
	start := time.Now()
	defer trackTime("Strategies Analysis", start)
	strategies := []string{
		"Chande Forecast Oscillator Strategy",
		"KDJ Strategy",
		"MACD Strategy",
		"Trend Stategy",
		"Volume Weighted Moving Average",
		"Awesome Oscillator Strategy",
		"RSI Strategy",
		"RSI 2 Strategy",
		"Williams R Strategy",
		"Bollinger Bands Strategy",
		"Chaikin Money Flow Strategy",
		"Ease of Movement Strategy",
		"Force Index Strategy",
		"Money Flow Index Strategy",
		"Negative Volume Index Strategy",
		"Volume Weighted Average Price Strategy",
	}

	actions := indicator.RunStrategies(
		asset,
		indicator.ChandeForecastOscillatorStrategy,
		indicator.DefaultKdjStrategy,
		indicator.MacdStrategy,
		indicator.MakeTrendStrategy(TREND_PERIOD),
		indicator.MakeVwmaStrategy(period),
		indicator.AwesomeOscillatorStrategy,
		indicator.DefaultRsiStrategy,
		indicator.Rsi2Strategy,
		indicator.WilliamsRStrategy,
		indicator.BollingerBandsStrategy,
		indicator.ChaikinMoneyFlowStrategy,
		indicator.EaseOfMovementStrategy,
		indicator.ForceIndexStrategy,
		indicator.MoneyFlowIndexStrategy,
		indicator.NegativeVolumeIndexStrategy,
		indicator.VolumeWeightedAveragePriceStrategy,
	)

	for index, stratActions := range actions {
		gains := indicator.ApplyActions(asset.Closing, stratActions)
		lastAction := stratActions[len(stratActions)-1]
		res := fmt.Sprintf("%v:: ", strategies[index])

		if lastAction == indicator.BUY {
			res += fmt.Sprintf("BUY recommended. ")
		} else if lastAction == indicator.SELL {
			res += fmt.Sprintf("SELL recommended. ")
		} else {
			res += fmt.Sprintf("HODL recommended. ")
		}

		res += fmt.Sprintf("GAINS = %.4f", gains[len(gains)-1])
		log.Info("Result", "data", res)
	}
}
