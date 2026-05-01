package vo

import (
	errs "dddcarpark/internal/errors"
)

type Mileage struct {
	distance uint32
	measure  Measure
}

func NewMileage(distance uint32, measure Measure) (Mileage, error) {
	if distance > 10_000_000 {
		return Mileage{}, errs.ErrInvalidDistance
	}
	return Mileage{
		distance: distance,
		measure:  measure,
	}, nil
}
