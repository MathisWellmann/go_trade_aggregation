package go_trade_aggregation

import (
	"errors"
	"fmt"
)

type (
	Scale struct {
		FromMin float64
		FromMax float64
		ToMin   float64
		ToMax   float64
	}
)

func NewScale(fromMin float64, fromMax float64, toMin float64, toMax float64) (*Scale, error) {
	if fromMax-fromMin == 0 {
		nErr := errors.New("fromMax - fromMin == 0")
		return nil, nErr
	}
	return &Scale{
		FromMin: fromMin,
		FromMax: fromMax,
		ToMin:   toMin,
		ToMax:   toMax,
	}, nil
}

func (s *Scale) Scale(val float64) (float64, error) {
	if val > s.FromMax {
		nErr := fmt.Sprintf("val > s.FromMax: %f > %f", val, s.FromMax)
		return 0, errors.New(nErr)
	} else if val < s.FromMin {
		nErr := fmt.Sprintf("val < s.FromMax: %f < %f", val, s.FromMin)
		return 0, errors.New(nErr)
	}
	out := s.ToMin + (((val - s.FromMin) * (s.ToMax - s.ToMin)) / (s.FromMax - s.FromMin))
	return out, nil
}
