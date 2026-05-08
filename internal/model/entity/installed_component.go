package entity

import (
	vo "dddcarpark/internal/model/valueobject"
	"time"

	"github.com/gofrs/uuid"
)

type InstalledComponent struct {
	ID               string
	ComponentNumber  string
	Type             vo.ComponentType
	InstalledAt      time.Time
	InstalledMileage vo.Mileage
	Resource         vo.Resource
}

func NewInstalledComponent(
	ComponentNumber string,
	Type vo.ComponentType,
	InstalledAt time.Time,
	currentMilleage vo.Mileage,
	resource vo.Resource,
) (InstalledComponent, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return InstalledComponent{}, err
	}
	return InstalledComponent{
		ID:               uuid.String(),
		ComponentNumber:  ComponentNumber,
		Type:             Type,
		InstalledAt:      InstalledAt,
		InstalledMileage: currentMilleage,
		Resource:         resource,
	}, nil
}
