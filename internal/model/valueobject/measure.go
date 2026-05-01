package vo

import errs "dddcarpark/internal/errors"

type Measure string

const (
	km    Measure = "km"
	miles Measure = "miles"
)

func NewMeasure(value string) (Measure, error) {
	measure := Measure(value)
	switch measure {
	case km, miles:
		return measure, nil
	default:
		return "", errs.ErrInvalidMeasure
	}
}
