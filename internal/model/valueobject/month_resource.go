package valueobject

import errs "dddcarpark/internal/errors"

type MonthResource uint32

func NewMonthResource(value uint32) (MonthResource, error) {
	if value < 1 || value > 1200 {
		return 0, errs.ErrInvalidMonthResource
	}
	return MonthResource(value), nil
}
