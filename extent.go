package go_trade_aggregation

import "math"

// ExtentCandles returns the min and max price over all candles
func ExtentCandles(candles []*Candle) (float64, float64) {
	high := candles[0].High
	low := candles[0].Low

	for i := 0; i < len(candles); i++ {
		if candles[i].High > high {
			high = candles[i].High
		}
		if candles[i].Low < low {
			low = candles[i].Low
		}
	}
	return low, high
}

// ExtentArr returns the Min and Max values of an array
func ExtentArr(vals []float64) []float64 {
	min := math.MaxFloat64
	max := math.SmallestNonzeroFloat64

	for i := 0; i < len(vals); i++ {
		if vals[i] > max {
			max = vals[i]
		}
		if vals[i] < min {
			min = vals[i]
		}
	}

	return []float64{min, max}
}
