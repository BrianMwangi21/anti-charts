package analysis

import (
	"fmt"
	"time"

	"github.com/cinar/indicator"
)

const TREND_PERIOD = 5

func performMACDStrategy(asset *indicator.Asset) string {
	start := time.Now()
	defer trackTime("Moving Average Convergence Divergence Strategy", start)
	res := "MACD STRATEGY:: TREND :: "
	actions := indicator.MacdStrategy(asset)
	actionsLen := len(actions)

	if actionsLen > 0 {
		lastAction := actions[actionsLen-1]

		if lastAction == indicator.BUY {
			res += fmt.Sprintf("BUY recommended. Action = %v", lastAction)
		} else if lastAction == indicator.SELL {
			res += fmt.Sprintf("SELL recommended. Action = %v", lastAction)
		} else {
			res += fmt.Sprintf("HODL recommended. Action = %v", lastAction)
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performTrendStrategy(asset *indicator.Asset) string {
	start := time.Now()
	defer trackTime("Trend Strategy", start)
	res := "TREND STRATEGY:: TREND :: "
	actions := indicator.TrendStrategy(asset, TREND_PERIOD)
	actionsLen := len(actions)

	if actionsLen > 0 {
		lastAction := actions[actionsLen-1]

		if lastAction == indicator.BUY {
			res += fmt.Sprintf("BUY recommended. Action = %v", lastAction)
		} else if lastAction == indicator.SELL {
			res += fmt.Sprintf("SELL recommended. Action = %v", lastAction)
		} else {
			res += fmt.Sprintf("HODL recommended. Action = %v", lastAction)
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performRSIStrategy(asset *indicator.Asset) string {
	start := time.Now()
	defer trackTime("Relatve Strength Index Strategy", start)
	res := "RSI STRATEGY:: MOMENTUN :: "
	actions := indicator.DefaultRsiStrategy(asset)
	actionsLen := len(actions)

	if actionsLen > 0 {
		lastAction := actions[actionsLen-1]

		if lastAction == indicator.BUY {
			res += fmt.Sprintf("BUY recommended. Action = %v", lastAction)
		} else if lastAction == indicator.SELL {
			res += fmt.Sprintf("SELL recommended. Action = %v", lastAction)
		} else {
			res += fmt.Sprintf("HODL recommended. Action = %v", lastAction)
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performBBStrategy(asset *indicator.Asset) string {
	start := time.Now()
	defer trackTime("Bollinger Bands Strategy", start)
	res := "BOLLINGER BANDS STRATEGY:: VOLATILITY :: "
	actions := indicator.BollingerBandsStrategy(asset)
	actionsLen := len(actions)

	if actionsLen > 0 {
		lastAction := actions[actionsLen-1]

		if lastAction == indicator.BUY {
			res += fmt.Sprintf("BUY recommended. Action = %v", lastAction)
		} else if lastAction == indicator.SELL {
			res += fmt.Sprintf("SELL recommended. Action = %v", lastAction)
		} else {
			res += fmt.Sprintf("HODL recommended. Action = %v", lastAction)
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performMFIStrategy(asset *indicator.Asset) string {
	start := time.Now()
	defer trackTime("Money Flow Index Strategy", start)
	res := "MFI STRATEGY:: VOLUME :: "
	actions := indicator.MoneyFlowIndexStrategy(asset)
	actionsLen := len(actions)

	if actionsLen > 0 {
		lastAction := actions[actionsLen-1]

		if lastAction == indicator.BUY {
			res += fmt.Sprintf("BUY recommended. Action = %v", lastAction)
		} else if lastAction == indicator.SELL {
			res += fmt.Sprintf("SELL recommended. Action = %v", lastAction)
		} else {
			res += fmt.Sprintf("HODL recommended. Action = %v", lastAction)
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performVWAPStrategy(asset *indicator.Asset) string {
	start := time.Now()
	defer trackTime("Volume Weighted Average Price Strategy", start)
	res := "VWAP STRATEGY:: VOLUME :: "
	actions := indicator.VolumeWeightedAveragePriceStrategy(asset)
	actionsLen := len(actions)

	if actionsLen > 0 {
		lastAction := actions[actionsLen-1]

		if lastAction == indicator.BUY {
			res += fmt.Sprintf("BUY recommended. Action = %v", lastAction)
		} else if lastAction == indicator.SELL {
			res += fmt.Sprintf("SELL recommended. Action = %v", lastAction)
		} else {
			res += fmt.Sprintf("HODL recommended. Action = %v", lastAction)
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performCumulativeStrategy(asset *indicator.Asset) string {
	start := time.Now()
	defer trackTime("Cumulative Strategy", start)
	res := "CUMULATIVE STRATEGY:: "

	strategies := indicator.AllStrategies(
		indicator.MacdStrategy,
		indicator.MakeTrendStrategy(TREND_PERIOD),
		indicator.DefaultRsiStrategy,
		indicator.BollingerBandsStrategy,
		indicator.MoneyFlowIndexStrategy,
		indicator.VolumeWeightedAveragePriceStrategy,
	)
	actions := strategies(asset)
	actionsLen := len(actions)

	if actionsLen > 0 {
		lastAction := actions[actionsLen-1]

		if lastAction == indicator.BUY {
			res += fmt.Sprintf("BUY recommended. Action = %v", lastAction)
		} else if lastAction == indicator.SELL {
			res += fmt.Sprintf("SELL recommended. Action = %v", lastAction)
		} else {
			res += fmt.Sprintf("HODL recommended. Action = %v", lastAction)
		}
	} else {
		res += "Insufficient data"
	}

	return res
}
