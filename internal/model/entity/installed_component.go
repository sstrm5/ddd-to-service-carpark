package entity

import (
	vo "dddcarpark/internal/model/valueobject"
	"time"
)

type InstalledComponent struct {
	ID               string
	ComponentNumber  string
	Type             vo.ComponentType
	InstalledAt      time.Time
	InstalledMileage vo.Mileage
	Resource         vo.Resource
}
