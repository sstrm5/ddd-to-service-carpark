package vo

import (
	errs "dddcarpark/internal/errors"
	"time"
)

type Year uint32

func NewYear(value uint32) (Year, error) {
	if value < 1900 || value > uint32(time.Now().Year()) {
		return 0, errs.ErrInvalidYear
	}
	return NewYear(value)
}
