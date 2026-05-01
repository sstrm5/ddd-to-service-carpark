package vo

import errs "dddcarpark/internal/errors"

type StatusOrder string

const (
	Draft     StatusOrder = "draft"
	Completed StatusOrder = "completed"
)

func NewStatusOrder(value string) (StatusOrder, error) {
	status := StatusOrder(value)
	switch status {
	case Draft, Completed:
		return status, nil
	default:
		return "", errs.ErrInvalidStatus
	}

}
