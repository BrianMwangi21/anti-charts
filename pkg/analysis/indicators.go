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
	res := "MACD:: "
	macd, signal := indicator.Macd(closes)
	macdLen, signalLen := len(macd), len(signal)

	if (macdLen > 0 && signalLen > 0) && (macdLen == signalLen) {
		aboveCount, belowCount, neutralCount := 0, 0, 0

		for index, macdLine := range macd {
			signalLine := signal[index]

			if macdLine > signalLine {
				aboveCount += 1
			} else if macdLine < signalLine {
				belowCount += 1
			} else {
				neutralCount += 1
			}
		}

		lastMacd := macd[macdLen-1]
		res += fmt.Sprintf("MACD > Signal Line %d times, MACD < Singal Line %d times, Neutral %d times. Last MACD = %.2f", aboveCount, belowCount, neutralCount, lastMacd)
	} else {
		res += "Insufficient data"
	}

	return res
}

func performSMA(closes []float64, period int, latestPrice float64) string {
	start := time.Now()
	defer trackTime("Simple Moving Average", start)
	res := "SMA:: "
	sma := indicator.Sma(period, closes)
	smaLen := len(sma)

	if smaLen > 0 {
		aboveCount, belowCount, neutralCount := 0, 0, 0

		for _, value := range sma {
			if value > latestPrice {
				aboveCount += 1
			} else if value < latestPrice {
				belowCount += 1
			} else {
				neutralCount += 1
			}
		}

		lastSma := sma[smaLen-1]
		res += fmt.Sprintf("SMA > Latest Price %d times, SMA < Latest Price %d times, Neutral %d times. Last SMA = %.2f, Latest Price = %.2f", aboveCount, belowCount, neutralCount, lastSma, latestPrice)
	} else {
		res += "Insufficient data"
	}

	return res
}

func performEMA(closes []float64, period int, latestPrice float64) string {
	start := time.Now()
	defer trackTime("Exponential Moving Average", start)
	res := "EMA:: "
	ema := indicator.Ema(period, closes)
	emaLen := len(ema)

	if emaLen > 0 {
		aboveCount, belowCount, neutralCount := 0, 0, 0

		for _, value := range ema {
			if value > latestPrice {
				aboveCount += 1
			} else if value < latestPrice {
				belowCount += 1
			} else {
				neutralCount += 1
			}
		}

		lastEma := ema[emaLen-1]
		res += fmt.Sprintf("EMA > Latest Price %d times, EMA < Latest Price %d times, Neutral %d times. Last EMA = %.2f, Latest Price = %.2f", aboveCount, belowCount, neutralCount, lastEma, latestPrice)
	} else {
		res += "Insufficient data"
	}

	return res
}

func performDEMA(closes []float64, period int, latestPrice float64) string {
	start := time.Now()
	defer trackTime("Double Exponential Moving Average", start)
	res := "DEMA:: "
	dema := indicator.Dema(period, closes)
	demaLen := len(dema)

	if demaLen > 0 {
		aboveCount, belowCount, neutralCount := 0, 0, 0

		for _, value := range dema {
			if value > latestPrice {
				aboveCount += 1
			} else if value < latestPrice {
				belowCount += 1
			} else {
				neutralCount += 1
			}
		}

		lastDema := dema[demaLen-1]
		res += fmt.Sprintf("DEMA > Latest Price %d times, DEMA < Latest Price %d times, Neutral %d times. Last DEMA = %.2f, Latest Price = %.2f", aboveCount, belowCount, neutralCount, lastDema, latestPrice)
	} else {
		res += "Insufficient data"
	}

	return res
}

func performTEMA(closes []float64, period int, latestPrice float64) string {
	start := time.Now()
	defer trackTime("Triple Exponential Moving Average", start)
	res := "TEMA:: "
	tema := indicator.Tema(period, closes)
	temaLen := len(tema)

	if temaLen > 0 {
		aboveCount, belowCount, neutralCount := 0, 0, 0

		for _, value := range tema {
			if value > latestPrice {
				aboveCount += 1
			} else if value < latestPrice {
				belowCount += 1
			} else {
				neutralCount += 1
			}
		}

		lastTema := tema[temaLen-1]
		res += fmt.Sprintf("TEMA above Latest Price %d times, TEMA below Latest Price %d times, Neutral %d times. Last TEMA = %.2f, Latest Price = %.2f", aboveCount, belowCount, neutralCount, lastTema, latestPrice)
	} else {
		res += "Insufficient data"
	}

	return res
}

func performRSI(closes []float64, latestPrice float64) string {
	start := time.Now()
	defer trackTime("Relative Strength Index", start)
	res := "RSI:: "
	rs, rsi := indicator.Rsi(closes)
	rsLen, rsiLen := len(rs), len(rsi)

	if (rsLen > 0 && rsiLen > 0) && (rsLen == rsiLen) {
		aboveCount, belowCount, neutralCount := 0, 0, 0

		for _, value := range rsi {
			if value > 70 {
				aboveCount += 1
			} else if value < 30 {
				belowCount += 1
			} else {
				neutralCount += 1
			}
		}

		lastRs := rs[rsLen-1]
		lastRsi := rsi[rsiLen-1]
		res += fmt.Sprintf("Overbought %d times, Oversold %d times, Neutral %d times. Last RS = %.2f, Last RSI = %.2f Latest Price = %.2f", aboveCount, belowCount, neutralCount, lastRs, lastRsi, latestPrice)
	} else {
		res += "Insufficient data"
	}

	return res
}

func performMFI(period int, highs, lows, closes, volumes []float64) string {
	start := time.Now()
	defer trackTime("Money Flow Index", start)
	res := "MFI:: "
	mfi := indicator.MoneyFlowIndex(period, highs, lows, closes, volumes)
	mfiLen := len(mfi)

	if mfiLen > 0 {
		aboveCount, belowCount, neutralCount := 0, 0, 0

		for i := 1; i < len(mfi); i++ {
			if mfi[i] > mfi[i-1] {
				aboveCount += 1
			} else if mfi[i] < mfi[i-1] {
				belowCount += 1
			} else {
				neutralCount += 1
			}
		}

		secondLastMfi := mfi[mfiLen-2]
		lastMfi := mfi[mfiLen-1]
		res += fmt.Sprintf("Uptrend %d times, Downtrend %d times, Neutral %d times. Second Last MFI = %.2f, Last MFI = %.2f", aboveCount, belowCount, neutralCount, secondLastMfi, lastMfi)
	} else {
		res += "Insufficient data"
	}

	return res
}

func trackTime(analysisName string, start time.Time) {
	elapsed := time.Since(start)
	log.Info(fmt.Sprintf("Performing %v", analysisName), "time elapsed", elapsed)
}
