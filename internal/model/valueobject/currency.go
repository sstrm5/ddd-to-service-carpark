package vo

import errs "dddcarpark/internal/errors"

type Currency string

const (
	Rub Currency = "rub"
	Usd Currency = "usd"
)

var validCurrencies = map[Currency]struct{}{
	Rub: {},
	Usd: {},
}

func NewCurrency(value string) (Currency, error) {
	currency := Currency(value)
	if _, ok := validCurrencies[currency]; !ok {
		return "", errs.ErrInvalidCurrency
	}
	return currency, nil
}
