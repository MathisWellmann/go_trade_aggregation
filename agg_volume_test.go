package go_trade_aggregation

import (
	"fmt"
	"testing"
)

func TestAggVolume(t *testing.T) {
	trades, err := LoadTradesFromCSV("data/Bitmex_XBTM20.csv")
	if err != nil {
		t.Error(err)
	}
	candles := AggVolume(trades, 700000)
	for i := 0; i < len(candles); i++ {
		err := errCheckAll(candles[i])
		if err {
			nErr := fmt.Sprintf("errCheckAll with candle: %#v", candles[i])
			t.Error(nErr)
		}
	}
}

func TestAggVolumeGraph(t *testing.T) {
	trades, err := LoadTradesFromCSV("data/Bitmex_XBTM20.csv")
	if err != nil {
		t.Error(err)
	}
	threshold := 4290000.0
	last := len(trades) - 1
	// aggregate last million trades
	candles := AggVolume(trades[last-1000000:last], threshold)

	fmt.Printf("len(candles): %d\n", len(candles))

	filename := fmt.Sprintf("img/agg_volume_%f.png", threshold)
	err = PlotCandleStick(candles, filename)
	if err != nil {
		t.Error(err)
	}
}
