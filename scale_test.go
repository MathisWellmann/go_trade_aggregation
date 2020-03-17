package go_trade_aggregation

import (
	"github.com/MathisWellmann/GoTrader/common/ml"
	"log"
	"math"
	"math/rand"
	"testing"
)

func TestScale(t *testing.T) {
	toMin := -1.0
	toMax := 1.0

	vals := genSymPos(128)
	extent := ExtentArr(vals)
	scale, err := NewScale(extent[0], extent[1], toMin, toMax)
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < len(vals); i++ {
		sVal, err := scale.Scale(vals[i])
		if err != nil {
			t.Error(err)
		}
		log.Printf("val: %f, sVal: %f", vals[i], sVal)
		if sVal > toMax {
			t.Error("sVal > toMax")
		} else if sVal < toMin {
			t.Error("sVal < toMin")
		}
	}
}

func genSymPos(length int) []float64 {
	out := make([]float64, length)
	out[0] = rand.Float64()
	for i := 1; i < len(out); i++ {
		change := ml.RandWeight()
		if out[i-1]+change <= 0 {
			change = math.Abs(change)
		}
		out[i] = out[i-1] + change
	}
	return out
}
