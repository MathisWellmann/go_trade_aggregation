package main

import (
	"fmt"
	. "github.com/MathisWellmann/go_trade_aggregation"
	"github.com/sirupsen/logrus"
)

func main() {
	trades, err := LoadTradesFromCSV("../data/Bitmex_XBTM20.csv")
	if err != nil {
		logrus.Fatal(err)
	}
	agg_volume := AggVolume(trades, 100)
	// print first 10 candles
	for i := 0; i < 10; i++ {
		fmt.Printf("agg_volume candle: %v", agg_volume[i])
	}

	agg_time := AggTime(trades, H1)
	// print first 10 candles
	for i := 0; i < 10; i++ {
		fmt.Printf("agg_time candle: %v", agg_time[i])
	}

	agg_market_energy := AggMarketEnergy(trades, 100)
	// print first 10 candles
	for i := 0; i < 10; i++ {
		fmt.Printf("agg_market_energy candle: %v", agg_market_energy[i])
	}
}
