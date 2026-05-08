package vo

import errs "dddcarpark/internal/errors"

type StatusVehicle string

const (
	Active      StatusVehicle = "active"       // машина готова к приему сервисом
	UnderRepair StatusVehicle = "under repair" // машина в сервисе
	WrittenOff  StatusVehicle = "written off"  // машина списана (больше не может проходить ТО)
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
