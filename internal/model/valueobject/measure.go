package vo

import errs "dddcarpark/internal/errors"

type Measure string

const (
	Km    Measure = "km"
	Miles Measure = "miles"
)

func NewMeasure(value string) (Measure, error) {
	measure := Measure(value)
	switch measure {
	case Km, Miles:
		return measure, nil
	default:
		return "", errs.ErrInvalidMeasure
	}
}
