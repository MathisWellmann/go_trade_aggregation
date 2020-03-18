package go_trade_aggregation

import (
	"encoding/csv"
	"fmt"
	"os"
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

func TestAggTimeExport(t *testing.T) {
	trades, err := LoadTradesFromCSV("data/Bitmex_XBTM20.csv")
	if err != nil {
		t.Error(err)
	}
	last := len(trades) - 1
	// aggregate last million trades
	candles := AggTime(trades[last-1000000:last], H1)

	fmt.Printf("len(candles): %d\n", len(candles))

	// export candles to csv for external plotting or analysis
	file, err := os.Create("data/agg_time.csv")
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
