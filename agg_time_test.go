package go_trade_aggregation

import (
	"fmt"
	"testing"
)

func TestAggTime(t *testing.T) {
	trades, err := LoadTradesFromCSV("data/Bitmex_XBTM20.csv")
	if err != nil {
		t.Error(err)
	}
	candles := AggTime(trades, H1)
	for i := 0; i < len(candles); i++ {
		err := errCheckAll(candles[i])
		if err {
			nErr := fmt.Sprintf("errCheckAll with candle: %#v", candles[i])
			t.Error(nErr)
		}
	}
}

func TestAggTimeGraph(t *testing.T) {
	trades, err := LoadTradesFromCSV("data/Bitmex_XBTM20.csv")
	if err != nil {
		t.Error(err)
	}
	last := len(trades) - 1
	candles := AggTime(trades[last-1000000:last], H1)

	fmt.Printf("len(candles): %d\n", len(candles))

	filename := "img/agg_time_h1.png"
	err = PlotCandleStick(candles, filename)
	if err != nil {
		t.Error(err)
	}
}
