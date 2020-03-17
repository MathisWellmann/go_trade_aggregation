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

func TestAggMarketEnergyGraph(t *testing.T) {
	trades, err := LoadTradesFromCSV("data/Bitmex_XBTM20.csv")
	if err != nil {
		t.Error(err)
	}
	last := len(trades) - 1
	threshold := 1585000.0
	candles := AggMarketEnergy(trades[last-100000:last], threshold)

	fmt.Printf("len(candles): %d\n", len(candles))

	filename := fmt.Sprintf("img/agg_market_energy_%d.png", int(threshold))
	err = PlotCandleStick(candles, filename)
	if err != nil {
		t.Error(err)
	}
}
