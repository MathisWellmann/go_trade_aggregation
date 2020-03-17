package main

import(
	"fmt"
	. "github.com/MathisWellmann/go_trade_aggregation"
)

func main() {
	trades := LoadTradesFromCSV("../data/XBTM20.csv")
	agg_volume := AggVolume(trades, 100)
	fmt.Printf("agg_volume: %v", agg_volume)
}
