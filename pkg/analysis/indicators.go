package analysis

import (
	"github.com/charmbracelet/log"
	"github.com/cinar/indicator"
)

func performMACD(closes []float64) string {
	log.Info("Performing Moving Average Convergence Divergence...")
	res := "MACD:: "
	macd, signal := indicator.Macd(closes)

	if len(macd) > 0 && len(signal) > 0 {
		macdLine := macd[len(macd)-1]
		signalLine := signal[len(signal)-1]

		if macdLine > signalLine {
			res += "MACD is above Signal Line"
		} else if macdLine < signalLine {
			res += "MACD is below Signal Line"
		} else {
			res += "MACD is equal to Signal Line"
		}

		if macdLine > 0 {
			res += ", MACD above 0"
		} else {
			res += ", MACD below 0"
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performSMA(closes []float64, period int, latestPrice float64) string {
	log.Info("Performing Simple Moving Average...")
	res := "SMA:: "
	sma := indicator.Sma(period, closes)
	smaLen := len(sma)

	if smaLen > 0 {
		lastSma := sma[smaLen-1]

		if latestPrice > lastSma {
			res += "Uptrend - Latest Price is above SMA"
		} else if latestPrice < lastSma {
			res += "Downtrend - Latest Price is below SMA"
		} else {
			res += "Neutral - Latest Price is at SMA level"
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performEMA(closes []float64, period int, latestPrice float64) string {
	log.Info("Performing Exponential Moving Average...")
	res := "EMA:: "
	ema := indicator.Ema(period, closes)
	emaLen := len(ema)

	if emaLen > 0 {
		lastEma := ema[emaLen-1]

		if latestPrice > lastEma {
			res += "Uptrend - Latest Price is above EMA"
		} else if latestPrice < lastEma {
			res += "Downtrend - Latest Price is below EMA"
		} else {
			res += "Neutral - Latest Price is at EMA level"
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performDEMA(closes []float64, period int, latestPrice float64) string {
	log.Info("Performing Double Exponential Moving Average...")
	dema := indicator.Dema(period, closes)
	demaLen := len(dema)
	res := "DEMA:: "

	if demaLen > 0 {
		lastDema := dema[demaLen-1]

		if latestPrice > lastDema {
			res += "Uptrend - Latest Price is above DEMA"
		} else if latestPrice < lastDema {
			res += "Downtrend - Latest Price is below DEMA"
		} else {
			res += "Neutral - Latest Price is at DEMA level"
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performTEMA(closes []float64, period int, latestPrice float64) string {
	log.Info("Performing Triple Exponential Moving Average...")
	tema := indicator.Tema(period, closes)
	temaLen := len(tema)
	res := "TEMA:: "

	if temaLen > 0 {
		lastTema := tema[temaLen-1]

		if latestPrice > lastTema {
			res += "Uptrend - Latest Price is above TEMA"
		} else if latestPrice < lastTema {
			res += "Downtrend - Latest Price is below TEMA"
		} else {
			res += "Neutral - Latest Price is at TEMA level"
		}
	} else {
		res += "Insufficient data"
	}

	return res
}
