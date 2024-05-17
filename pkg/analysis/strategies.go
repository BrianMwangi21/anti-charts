package analysis

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/cinar/indicator"
)

const TREND_PERIOD = 5

type Strat struct {
	Name   string
	Weight int
	Result string
}

func performAllStrategies(asset *indicator.Asset, period int) {
	strats := []Strat{
		{"MACD Strategy", 5, ""},
		{"RSI Strategy", 5, ""},
		//
		{"Chande Forecast Oscillator Strategy", 4, ""},
		{"Trend Stategy", 4, ""},
		{"Bollinger Bands Strategy", 4, ""},
		{"Money Flow Index Strategy", 4, ""},
		{"Volume Weighted Average Price Strategy", 4, ""},
		//
		{"KDJ Strategy", 3, ""},
		{"Volume Weighted Moving Average", 3, ""},
		{"Awesome Oscillator Strategy", 3, ""},
		{"Williams R Strategy", 3, ""},
		{"Chaikin Money Flow Strategy", 3, ""},
		{"Ease of Movement Strategy", 3, ""},
		{"Force Index Strategy", 3, ""},
		//
		{"Negative Volume Index Strategy", 2, ""},
		{"RSI 2 Strategy", 2, ""},
	}

	strategies := []indicator.StrategyFunction{
		indicator.MacdStrategy,
		indicator.DefaultRsiStrategy,
		//
		indicator.ChandeForecastOscillatorStrategy,
		indicator.MakeTrendStrategy(TREND_PERIOD),
		indicator.BollingerBandsStrategy,
		indicator.MoneyFlowIndexStrategy,
		indicator.VolumeWeightedAveragePriceStrategy,
		//
		indicator.DefaultKdjStrategy,
		indicator.MakeVwmaStrategy(period),
		indicator.AwesomeOscillatorStrategy,
		indicator.WilliamsRStrategy,
		indicator.ChaikinMoneyFlowStrategy,
		indicator.EaseOfMovementStrategy,
		indicator.ForceIndexStrategy,
		//
		indicator.NegativeVolumeIndexStrategy,
		indicator.Rsi2Strategy,
	}

	actions := indicator.RunStrategies(asset, strategies...)

	for index, stratActions := range actions {
		gains := indicator.ApplyActions(asset.Closing, stratActions)
		lastAction := stratActions[len(stratActions)-1]
		res := fmt.Sprintf("%v:: ", strats[index].Name)

		if lastAction == indicator.BUY {
			res += fmt.Sprintf("BUY recommended. ")
		} else if lastAction == indicator.SELL {
			res += fmt.Sprintf("SELL recommended. ")
		} else {
			res += fmt.Sprintf("HOLD recommended. ")
		}

		res += fmt.Sprintf("Gains = %.4f", gains[len(gains)-1])
		strats[index].Result = res
	}

	log.Info("Strategies Results By Weight...")
	currentWeight := 6
	for _, value := range strats {
		if value.Weight != currentWeight {
			currentWeight -= 1
			log.Info(fmt.Sprintf("=== Strategy Weight %d ===", currentWeight))
		}
		log.Info("STRATEGIES", "result", value.Result)
	}
}
