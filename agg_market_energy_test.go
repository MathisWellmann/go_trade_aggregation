package go_trade_aggregation

import (
	"fmt"
	"testing"
)

func TestAggMarketEnergy(t *testing.T) {
	trades, err := LoadTradesFromCSV("data/Bitmex_XBTM20.csv")
	if err != nil {
		t.Error(err)
	}
	candles := AggMarketEnergy(trades, 100)
	for i := 0; i < len(candles); i++ {
		err := errCheckAll(candles[i])
		if err {
			nErr := fmt.Sprintf("errCheckAll with candle: %#v", candles[i])
			t.Error(nErr)
		}
	}
}
