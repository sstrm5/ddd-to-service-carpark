package vo

import (
	errs "dddcarpark/internal/errors"
)

type Mileage struct {
	distance uint32
	measure  Measure
}

const MilesToKmCoeff float32 = 1.61

func NewMileage(distance uint32, measure Measure) (Mileage, error) {
	if distance > 10_000_000 {
		return Mileage{}, errs.ErrInvalidDistance
	}
	return Mileage{
		distance: distance,
		measure:  measure,
	}, nil
}

func (m *Mileage) LessThan(other Mileage) bool {
	if m.measure != other.measure {
		if m.measure == Miles {
			return float32(m.distance)*MilesToKmCoeff < float32(other.distance)
		}
		if other.measure == Miles {
			return float32(m.distance) < float32(other.distance)*MilesToKmCoeff
		}
	}
	return m.distance < other.distance
}
