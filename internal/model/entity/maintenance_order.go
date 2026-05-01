package entity

import (
	vo "dddcarpark/internal/model/valueobject"
	"time"
)

type MaintenanceOrder struct {
	ID             string
	VehicleID      string
	CreatedAt      time.Time
	ServiceMileage vo.Mileage
	Status         vo.StatusVehicle
	Currency       vo.Currency
}
