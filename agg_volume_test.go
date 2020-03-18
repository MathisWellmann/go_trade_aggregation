package go_trade_aggregation

import (
	"encoding/csv"
	"fmt"
	"os"
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

func TestAggVolumeExport(t *testing.T) {
	trades, err := LoadTradesFromCSV("data/Bitmex_XBTM20.csv")
	if err != nil {
		t.Error(err)
	}
	threshold := 4290000.0
	last := len(trades) - 1
	// aggregate last million trades
	candles := AggVolume(trades[last-1000000:last], threshold)

	fmt.Printf("len(candles): %d\n", len(candles))

	// export candles to csv for external plotting or analysis
	file, err := os.Create("data/agg_volume.csv")
	if err != nil {
		t.Error(err)
	}

	w := csv.NewWriter(file)
	// only export OHLCV candles, ignore other filed in Candle
	header := []string{
		"timestamp",
		"open",
		"high",
		"low",
		"close",
		"volume",
	}
	err = w.Write(header)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(candles); i++ {
		ts := fmt.Sprintf("%d", candles[i].Timestamp)
		open := fmt.Sprintf("%f", candles[i].Open)
		high := fmt.Sprintf("%f", candles[i].High)
		low := fmt.Sprintf("%f", candles[i].Low)
		close := fmt.Sprintf("%f", candles[i].Close)
		row := []string{
			ts,
			open,
			high,
			low,
			close,
		}
		err = w.Write(row)
		if err != nil {
			t.Error(err)
		}
	}
}
