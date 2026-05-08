package vo

import (
	errs "dddcarpark/internal/errors"
)

type Mileage struct {
	Distance uint32
	Measure  Measure
}

const MilesToKmCoeff float32 = 1.61

func NewMileage(distance uint32, measure Measure) (Mileage, error) {
	if distance > 10_000_000 {
		return Mileage{}, errs.ErrInvalidDistance
	}
	return Mileage{
		Distance: distance,
		Measure:  measure,
	}, nil
}

func (m *Mileage) LessThan(other Mileage) bool {
	if m.Measure != other.Measure {
		if m.Measure == Miles {
			return float32(m.Distance)*MilesToKmCoeff < float32(other.Distance)
		}
		if other.Measure == Miles {
			return float32(m.Distance) < float32(other.Distance)*MilesToKmCoeff
		}
	}
	return m.Distance < other.Distance
}
