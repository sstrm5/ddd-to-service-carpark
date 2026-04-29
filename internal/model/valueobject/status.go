package valueobject

import errs "dddcarpark/internal/errors"

type Status string

const (
	Active      Status = "active"
	UnderRepair Status = "under repair"
	WrittenOff  Status = "written off"
)

func NewStatus(value string) (Status, error) {
	status := Status(value)
	switch status {
	case Active, UnderRepair, WrittenOff:
		return status, nil
	default:
		return "", errs.ErrInvalidStatus
	}

}
