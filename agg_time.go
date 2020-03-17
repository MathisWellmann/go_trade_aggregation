package go_trade_aggregation

import "math"

const (
	M1  = 60
	M5  = 300
	M15 = 900
	M30 = 1800
	H1  = 3600
	H2  = 7200
	H4  = 14400
	H6  = 21600
	H8  = 28800
	H12 = 43200
	D1  = 86400
)

// AggTime aggregates trades by time and returns candles
// threshold is in seconds
func AggTime(trades []*Trade, threshold int64) []*Candle {
	var out []*Candle

	var init_ts int64
	var open float64
	var high float64
	var low float64
	var volume float64
	var buyVolume float64
	var numBuys int
	var numTrades int
	var wp float64 // used for calculating weighted price
	init := true

	for i := 0; i < len(trades); i++ {
		if init {
			init = false
			// set initial values for next candle
			init_ts = trades[i].Timestamp
			open = trades[i].Price
			high = trades[i].Price
			low = trades[i].Price
			volume = 0
			buyVolume = 0
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
			buyVolume += trades[i].Size
		}
		wp += trades[i].Price * math.Abs(trades[i].Size)

		// convert threshold value from seconds to milliseconds like the timestamps are
		if trades[i].Timestamp-init_ts > threshold*1000 {
			// create new candle
			c := &Candle{
				Timestamp:            trades[i].Timestamp,
				Open:                 open,
				High:                 high,
				Low:                  low,
				Close:                trades[i].Price,
				Volume:               volume,
				NumTrades:            numTrades,
				TradeDirectionRatio:  float64(numBuys) / float64(numTrades),
				VolumeDirectionRatio: buyVolume / volume,
				WeightedPrice:        wp / volume,
			}
			out = append(out, c)

			init = true
		}
	}

	return out
}
