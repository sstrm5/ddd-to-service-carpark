package vo

import errs "dddcarpark/internal/errors"

type StatusVehicle string

const (
	Active      StatusVehicle = "active"
	UnderRepair StatusVehicle = "under repair"
	WrittenOff  StatusVehicle = "written off"
)

func NewStatusVehicle(value string) (StatusVehicle, error) {
	status := StatusVehicle(value)
	switch status {
	case Active, UnderRepair, WrittenOff:
		return status, nil
	default:
		return "", errs.ErrInvalidStatus
	}

}
