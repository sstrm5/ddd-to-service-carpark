package vo

import errs "dddcarpark/internal/errors"

type Money struct {
	Value    uint32
	Currency Currency
}

func NewMoney(value uint32, currency Currency) (Money, error) {
	if currency == "" {
		return Money{}, errs.ErrInvalidCurrency
	}
	return Money{
		Value:    value,
		Currency: currency,
	}, nil
}
