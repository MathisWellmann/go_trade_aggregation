package go_trade_aggregation

import "math"

// AggVolume aggregates trades by volume and return the candles
func AggVolume(trades []*Trade, threshold float64) []*Candle {
	var out []*Candle

	var open float64
	var high float64
	var low float64
	var volume float64
	var numBuys int
	var numTrades int
	var wp float64 // used for calculating weighted price
	init := true

	for i := 0; i < len(trades); i++ {
		if init {
			init = false
			// set initial values for next candle
			open = trades[i].Price
			high = trades[i].Price
			low = trades[i].Price
			volume = 0
			numBuys = 0
			numTrades = 0
			wp = 0
		}

		if trades[i].Price > high {
			high = trades[i].Price
		} else if trades[i].Price < low {
			low = trades[i].Price
		}
		volume += math.Abs(trades[i].Size)
		numTrades++
		if trades[i].Size > 0 {
			numBuys++
		}
		wp += trades[i].Price * math.Abs(trades[i].Size)

		if volume > threshold {
			// create new candle
			c := &Candle{
				Timestamp:           trades[i].Timestamp,
				Open:                open,
				High:                high,
				Low:                 low,
				Close:               trades[i].Price,
				Volume:              volume,
				NumTrades:           numTrades,
				TradeDirectionRatio: float64(numBuys) / float64(numTrades),
				WeightedPrice:       wp / volume,
			}
			out = append(out, c)

			init = true
		}
	}

	return out
}

// TODO: helper function to convert from time period to volume period
