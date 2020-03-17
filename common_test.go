package go_trade_aggregation

// This file provides helper functions for testing

// errCheckAll will run all tests on candle
// returns true if any value in candle is invalid
func errCheckAll(candle *Candle) bool {
	err := errOpen(candle)
	if err {
		return true
	}
	err = errHigh(candle)
	if err {
		return true
	}
	err = errLow(candle)
	if err {
		return true
	}
	err = errClose(candle)
	if err {
		return true
	}
	err = errNumTrades(candle)
	if err {
		return true
	}
	err = errTradeDirectionRatio(candle)
	if err {
		return true
	}
	err = errWeightedPrice(candle)
	if err {
		return true
	}
	return false
}

// errOpen checks if the open price of candle is valid
// returns true if error is found
func errOpen(candle *Candle) bool {
	if candle.Open > candle.High {
		return true
	}
	if candle.Open < candle.Low {
		return true
	}
	return false
}

// errHigh checks if the high of candle is valid
// returns true if error is found
func errHigh(candle *Candle) bool {
	if candle.High < candle.Low {
		return true
	}
	if candle.High < candle.Open {
		return true
	}
	if candle.High < candle.Close {
		return true
	}
	return false
}

// errLow check is the low of candle is valid
// returns true if error is found
func errLow(candle *Candle) bool {
	if candle.Low > candle.Open {
		return true
	}
	if candle.Low > candle.High {
		return true
	}
	if candle.Low > candle.Close {
		return true
	}
	return false
}

// errClose checks if the close price of candle is valid
// return true if error has been found
func errClose(candle *Candle) bool {
	if candle.Close > candle.High {
		return true
	}
	if candle.Close < candle.Low {
		return true
	}
	return false
}

// errTradeDirectionRatio checks the correctness of TradeDirectionRatio within the candle
// returns true if error has been found
func errTradeDirectionRatio(candle *Candle) bool {
	if candle.TradeDirectionRatio > 1 {
		return true
	}
	if candle.TradeDirectionRatio < 0 {
		return true
	}
	return false
}

// errWeightedPrice checks if WeightedPrice of candle is withing candle range
// returns true if error is found
func errWeightedPrice(candle *Candle) bool {
	if candle.WeightedPrice > candle.High {
		return true
	}
	if candle.WeightedPrice < candle.Low {
		return true
	}
	return false
}

// errNumTrades checks if candle contains any trades at all
// returns true if err is found
func errNumTrades(candle *Candle) bool {
	if candle.NumTrades == 0 {
		return true
	}
	return false
}
