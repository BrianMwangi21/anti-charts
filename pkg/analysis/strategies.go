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
	Action indicator.Action
}

var STRATS = []Strat{
	{"MACD Strategy", 5, "", indicator.HOLD},
	{"RSI Strategy", 5, "", indicator.HOLD},
	//
	{"Chande Forecast Oscillator Strategy", 4, "", indicator.HOLD},
	{"Trend Stategy", 4, "", indicator.HOLD},
	{"Bollinger Bands Strategy", 4, "", indicator.HOLD},
	{"Money Flow Index Strategy", 4, "", indicator.HOLD},
	{"Volume Weighted Average Price Strategy", 4, "", indicator.HOLD},
	//
	{"KDJ Strategy", 3, "", indicator.HOLD},
	{"Volume Weighted Moving Average", 3, "", indicator.HOLD},
	{"Awesome Oscillator Strategy", 3, "", indicator.HOLD},
	{"Williams R Strategy", 3, "", indicator.HOLD},
	{"Chaikin Money Flow Strategy", 3, "", indicator.HOLD},
	{"Ease of Movement Strategy", 3, "", indicator.HOLD},
	{"Force Index Strategy", 3, "", indicator.HOLD},
	//
	{"Negative Volume Index Strategy", 2, "", indicator.HOLD},
	{"RSI 2 Strategy", 2, "", indicator.HOLD},
}

func performAllStrategies(asset *indicator.Asset, period int) indicator.Action {
	var buys, sells, holds int

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
		res := fmt.Sprintf("%v:: ", STRATS[index].Name)

		if lastAction == indicator.BUY {
			res += fmt.Sprintf("BUY recommended. ")
			buys += 1
		} else if lastAction == indicator.SELL {
			res += fmt.Sprintf("SELL recommended. ")
			sells += 1
		} else {
			res += fmt.Sprintf("HOLD recommended. ")
			holds += 1
		}

		res += fmt.Sprintf("Gains = %.4f", gains[len(gains)-1])
		STRATS[index].Result = res
		STRATS[index].Action = lastAction
	}

	totalN := float64(len(STRATS))
	buyP := (float64(buys) / totalN) * 100
	sellP := (float64(sells) / totalN) * 100
	holdP := (float64(holds) / totalN) * 100
	finalAction, aBuyP, aSellP, aHoldP := aggregateResults()

	log.Info("Strategies Results By Weight...")
	currentWeight := 6
	for _, strat := range STRATS {
		if strat.Weight != currentWeight {
			currentWeight -= 1
			log.Info(fmt.Sprintf("=== Strategy Weight %d ===", currentWeight))
		}
		log.Info("STRATEGIES", "result", strat.Result)
	}
	log.Info(fmt.Sprintf("STRATEGIES Normal Summary :: BUYS = %d [%.2f], SELLS = %d [%.2f], HOLDS = %d [%.2f]", buys, buyP, sells, sellP, holds, holdP))
	log.Info(fmt.Sprintf("STRATEGIES Weighted Summary :: BUYS = %d [%.2f], SELLS = %d [%.2f], HOLDS = %d [%.2f]", buys, aBuyP, sells, aSellP, holds, aHoldP))
	log.Info(fmt.Sprintf("STRATEGIES FINAL ACTION :: %v", finalAction))

	return finalAction
}

func aggregateResults() (indicator.Action, float64, float64, float64) {
	var buyW, sellW, holdW, totalW int

	for _, strat := range STRATS {
		totalW += strat.Weight
		switch strat.Action {
		case indicator.BUY:
			buyW += strat.Weight
		case indicator.SELL:
			sellW += strat.Weight
		case indicator.HOLD:
			holdW += strat.Weight
		}
	}

	totalP := float64(totalW)
	aBuyP := (float64(buyW) / totalP) * 100
	aSellP := (float64(sellW) / totalP) * 100
	aHoldP := (float64(holdW) / totalP) * 100

	if aBuyP > aSellP && aBuyP > aHoldP {
		return indicator.BUY, aBuyP, aSellP, aHoldP
	} else if aSellP > aBuyP && aSellP > aHoldP {
		return indicator.SELL, aBuyP, aSellP, aHoldP
	} else {
		return indicator.HOLD, aBuyP, aSellP, aHoldP
	}
}
