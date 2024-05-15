package analysis

import (
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	"github.com/cinar/indicator"
)

func performMACD(closes []float64) string {
	start := time.Now()
	defer trackTime("Moving Average Convergence Divergence", start)
	res := "MACD:: TREND :: "
	macd, signal := indicator.Macd(closes)
	macdLen, signalLen := len(macd), len(signal)

	if (macdLen > 0 && signalLen > 0) && (macdLen == signalLen) {
		histogram := make([]float64, macdLen)

		for i := range macd {
			histogram[i] = macd[i] - signal[i]
		}

		latestMACD := macd[macdLen-1]
		latestSignal := signal[signalLen-1]
		latestHistogram := histogram[macdLen-1]

		if latestMACD > latestSignal && latestHistogram > 0 {
			res += fmt.Sprintf("Bullish. MACD = %.2f, Signal = %.2f, Histogram = %.2f", latestMACD, latestSignal, latestHistogram)
		} else if latestMACD < latestSignal && latestHistogram < 0 {
			res += fmt.Sprintf("Bearish. MACD = %.2f, Signal = %.2f, Histogram = %.2f", latestMACD, latestSignal, latestHistogram)
		} else {
			res += fmt.Sprintf("Neutral. MACD = %.2f, Signal = %.2f, Histogram = %.2f", latestMACD, latestSignal, latestHistogram)
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performSMA(closes []float64, period int, latestPrice float64) string {
	start := time.Now()
	defer trackTime("Simple Moving Average", start)
	res := "SMA:: TREND :: "
	sma := indicator.Sma(period, closes)
	smaLen := len(sma)

	if smaLen > 0 {
		latestSMA := sma[smaLen-1]

		if latestPrice > latestSMA {
			res += fmt.Sprintf("Bullish. SMA = %.2f", latestSMA)
		} else if latestPrice < latestSMA {
			res += fmt.Sprintf("Bearish. SMA = %.2f", latestSMA)
		} else {
			res += fmt.Sprintf("Neutral. SMA = %.2f", latestSMA)
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performEMA(closes []float64, period int, latestPrice float64) string {
	start := time.Now()
	defer trackTime("Exponential Moving Average", start)
	res := "EMA:: TREND :: "
	ema := indicator.Ema(period, closes)
	emaLen := len(ema)

	if emaLen > 0 {
		latestEMA := ema[emaLen-1]

		if latestPrice > latestEMA {
			res += fmt.Sprintf("Bullish. EMA = %.2f", latestEMA)
		} else if latestPrice < latestEMA {
			res += fmt.Sprintf("Bearish. EMA = %.2f", latestEMA)
		} else {
			res += fmt.Sprintf("Neutral. EMA = %.2f", latestEMA)
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performDEMA(closes []float64, period int, latestPrice float64) string {
	start := time.Now()
	defer trackTime("Double Exponential Moving Average", start)
	res := "DEMA:: TREND :: "
	dema := indicator.Dema(period, closes)
	demaLen := len(dema)

	if demaLen > 0 {
		latestDEMA := dema[demaLen-1]

		if latestPrice > latestDEMA {
			res += fmt.Sprintf("Bullish. DEMA = %.2f", latestDEMA)
		} else if latestPrice < latestDEMA {
			res += fmt.Sprintf("Bearish. DEMA = %.2f", latestDEMA)
		} else {
			res += fmt.Sprintf("Neutral. DEMA = %.2f", latestDEMA)
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performTEMA(closes []float64, period int, latestPrice float64) string {
	start := time.Now()
	defer trackTime("Triple Exponential Moving Average", start)
	res := "TEMA:: TREND :: "
	tema := indicator.Tema(period, closes)
	temaLen := len(tema)

	if temaLen > 0 {
		latestTEMA := tema[temaLen-1]

		if latestPrice > latestTEMA {
			res += fmt.Sprintf("Bullish. TEMA = %.2f", latestTEMA)
		} else if latestPrice < latestTEMA {
			res += fmt.Sprintf("Bearish. TEMA = %.2f", latestTEMA)
		} else {
			res += fmt.Sprintf("Neutral. TEMA = %.2f", latestTEMA)
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performRSI(closes []float64) string {
	start := time.Now()
	defer trackTime("Relative Strength Index", start)
	res := "RSI:: MOMENTUM :: "
	rs, rsi := indicator.Rsi(closes)
	rsLen, rsiLen := len(rs), len(rsi)

	if (rsLen > 0 && rsiLen > 0) && (rsLen == rsiLen) {
		latestRSI := rsi[rsiLen-1]

		if latestRSI > 70 {
			res += fmt.Sprintf("Overbought. RSI = %.2f", latestRSI)
		} else if latestRSI < 30 {
			res += fmt.Sprintf("Oversold. RSI = %.2f", latestRSI)
		} else {
			res += fmt.Sprintf("Neutral. RSI = %.2f", latestRSI)
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performBB(closes []float64, latestPrice float64) string {
	start := time.Now()
	defer trackTime("Bollinger Bands", start)
	res := "BB:: VOLATILITY :: "
	middle, upper, lower := indicator.BollingerBands(closes)
	middleLen, upperLen, lowerLen := len(middle), len(upper), len(lower)

	if middleLen > 0 && upperLen > 0 && lowerLen > 0 {
		if latestPrice > upper[upperLen-1] {
			res += fmt.Sprintf("High Volatility - Price above upper band. %.2f", upper[upperLen-1])
		} else if latestPrice < lower[lowerLen-1] {
			res += fmt.Sprintf("High Volatility - Price below lower band. %.2f", lower[lowerLen-1])
		} else if latestPrice > middle[middleLen-1] {
			res += fmt.Sprintf("Medium Volatility - Price above middle band. %.2f", middle[middleLen-1])
		} else if latestPrice < middle[middleLen-1] {
			res += fmt.Sprintf("Medium Volatility - Price below middle band. %.2f", middle[middleLen-1])
		} else {
			res += fmt.Sprintf("Low Volatility - Price within Bollinger Bands")
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func performMFI(period int, highs, lows, closes, volumes []float64) string {
	start := time.Now()
	defer trackTime("Money Flow Index", start)
	res := "MFI:: VOLUME :: "
	mfi := indicator.MoneyFlowIndex(period, highs, lows, closes, volumes)
	mfiLen := len(mfi)

	if mfiLen > 0 {
		latestMFI := mfi[mfiLen-1]

		if latestMFI > 80 {
			res += fmt.Sprintf("Overbought. MFI = %.2f", latestMFI)
		} else if latestMFI < 20 {
			res += fmt.Sprintf("Oversold. MFI = %.2f", latestMFI)
		} else {
			res += fmt.Sprintf("Neutral. MFI = %.2f", latestMFI)
		}
	} else {
		res += "Insufficient data"
	}

	return res
}

func trackTime(analysisName string, start time.Time) {
	elapsed := time.Since(start)
	log.Info(fmt.Sprintf("Performing %v", analysisName), "time elapsed", elapsed)
}
