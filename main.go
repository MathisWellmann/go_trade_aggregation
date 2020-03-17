package go_trade_aggregation

type (
	Trade struct {
		Timestamp int64 // in Milliseconds
		Price     float64
		Size      float64 // will be negative if sell trade
	}
	Candle struct {
		Timestamp           int64
		Open                float64
		High                float64
		Low                 float64
		Close               float64
		Volume              float64
		NumTrades           int
		TradeDirectionRatio float64 // numBuys / numTrades
		WeightedPrice       float64
	}
)
