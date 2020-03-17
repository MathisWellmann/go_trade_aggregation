# Trade Aggregation
Convert trade data into candles using different forms of aggregation.
The Candles provide more detailed statistics than the usual OHLCV candles.
Additional statistics inlcude:
- number of trades
- trade direction ratio ( num_buys / num_trades )
- weighted average price ( using abs(size) as weight)

This Aggregation package allows for the creation of highly sophisticated algorithm and ML models.
It enables a clear view of the market state without using arbitrary time aggregation.

### Time Aggregation:
Creates a candle every n seconds.
This method of aggregating trades has a long history mostly due to humans interacting with the market with a perception of time.
This is however not how the market works (especially 24/7 crypto markets).
The markets doesnt care about the time, only about price and volume.

### Volume Aggregation:
Creates a candle every n traded contracts ( trade size)
Price moves occur whenever and aggressive trader places a market order with a given size.
The more size is being traded the more likely a stronger move will be.
This is more natural way to view the market and provides many advantages over time aggregation such as more well behaved volatility.
In this mode of Aggregation, candles will be printed dynamically and will print more candles in times of higher volume / volatility,
therefore providing the trader which an automatically scaling view, just like as if he would switch time periods, but way better.

### Market Energy Aggregation:
Creates candles using the aggregation function 
``
sqrt(abs(size) * abs(return))
``
This is the most sophisticated aggregation method and captures the market energy well, creating a timeseries which is well behaved and does not suffer as much as time aggregated candles from volatility.

### Images

### How To Use:
First load your desired trades into []*Trade. Note that if trade is sell, then size is negative. 
This reduces memory usage over storing that info in string as it is only one bit.
See example folder and tests for more details.

### TODOs:
- Helper functions for converting time period to other threshold values so that the same number of candles can be
returned over all aggregation methods without tuning parameter manually.
- Analysis paper of observed behaviour with different aggregation methods