package go_trade_aggregation

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/softwarebackend"
	"image/png"
	"os"
)

func PlotCandleStick(candles []*Candle, filename string) error {
	wpx := 1920
	hpx := 1080
	backend := softwarebackend.New(wpx, hpx)
	cv := canvas.New(backend)

	width, height := cv.Size()
	min, max := ExtentCandles(candles)
	xScale, _ := NewScale(0, 1, 0, float64(width))
	yScale, _ := NewScale(min, max, 0, float64(height))

	cv.SetFillStyle("#000")
	cv.FillRect(0, 0, float64(width), float64(height))

	for i := 0; i < len(candles); i++ {
		var candleColor string
		if candles[i].Close >= candles[i].Open {
			candleColor = "#00FF00"
			fmt.Printf("GREEN candle: %#v\n", candles[i])
		} else {
			candleColor = "#FF0000"
		}

		cv.SetFillStyle(candleColor)
		x, _ := xScale.Scale(float64(i) / float64(len(candles)))
		y, _ := yScale.Scale(candles[i].Low)
		candleWidth, _ := xScale.Scale(1.0 / float64(len(candles)))
		candleHeight := ((candles[i].High - candles[i].Low) / max) * float64(height)

		cv.FillRect(x, y, candleWidth, candleHeight)
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		logrus.Fatal(err)
	}
	defer file.Close()

	err = png.Encode(file, backend.Image)
	if err != nil {
		logrus.Error(err)
	}

	return nil
}
