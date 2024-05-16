package analysis

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/cinar/indicator"
)

const TREND_PERIOD = 5

func performAllStrategies(asset *indicator.Asset, period int) {
	var (
		buys  []string
		holds []string
		sells []string
	)
	strategyNames := []string{
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

	strategies := []indicator.StrategyFunction{
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
	}

	actions := indicator.RunStrategies(asset, strategies...)

	for index, stratActions := range actions {
		gains := indicator.ApplyActions(asset.Closing, stratActions)
		lastAction := stratActions[len(stratActions)-1]
		res := fmt.Sprintf("%v:: ", strategyNames[index])

		if lastAction == indicator.BUY {
			res += fmt.Sprintf("BUY recommended. Gains = %.4f", gains[len(gains)-1])
			buys = append(buys, res)
		} else if lastAction == indicator.SELL {
			res += fmt.Sprintf("SELL recommended. Gains = %.4f", gains[len(gains)-1])
			sells = append(sells, res)
		} else {
			res += fmt.Sprintf("HOLD recommended. Gains = %.4f", gains[len(gains)-1])
			holds = append(holds, res)
		}
	}

	log.Info("Strategies Recommending BUY...")
	for _, value := range buys {
		log.Info("STRATEGIES", "result", value)
	}

	log.Info("Strategies Recommending HODL...")
	for _, value := range holds {
		log.Info("STRATEGIES", "result", value)
	}

	log.Info("Strategies Recommending SELL...")
	for _, value := range sells {
		log.Info("STRATEGIES", "result", value)
	}
}
