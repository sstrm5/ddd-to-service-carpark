package entity

import (
	errs "dddcarpark/internal/errors"
	vo "dddcarpark/internal/model/valueobject"
)

type Vehicle struct {
	ID         string
	Vin        vo.VinNumber
	Make       string
	Model      string
	Mileage    vo.Mileage
	Year       vo.Year
	Status     vo.StatusVehicle
	Components map[vo.ComponentType]*InstalledComponent
}

func NewVehicle(
	vin vo.VinNumber,
	vehicleMake string,
	vehicleModel string,
	year vo.Year,
	initMileage vo.Mileage,
) (Vehicle, error) {
	if vehicleMake == "" {
		return Vehicle{}, errs.ErrInvalidVehicleMake
	}
	if vehicleModel == "" {
		return Vehicle{}, errs.ErrInvalidVehicleModel
	}
	return Vehicle{
		Vin:     vin,
		Make:    vehicleMake,
		Model:   vehicleModel,
		Mileage: initMileage,
		Year:    year,
		Status:  vo.Active,
	}, nil
}

func (v *Vehicle) ValidateForDraftOrder() error {
	switch v.Status {
	case vo.Active:
		return nil
	case vo.UnderRepair:
		return errs.ErrVehicleUnderRepair
	case vo.WrittenOff:
		return errs.ErrVehicleWritenOff
	default:
		return errs.ErrInvalidVehicleStatus
	}
}

// func (v *Vehicle) InstallComponent(
// 	spec vo.ComponentSpec,
// 	currentComponentMileage vo.Mileage,
// 	componentResource vo.Resource,
// ) (InstalledComponent, error) {
// 	component, err := NewInstalledComponent(
// 		spec.ComponentNumber,
// 		spec.Type,
// 		time.Now(),
// 		currentComponentMileage,
// 		componentResource,
// 	)
// 	if err != nil {
// 		return InstalledComponent{}, nil
// 	}
// 	v.Components[spec.Type] = &component
// 	return component, nil
// }

func (v *Vehicle) InstallComponent(component InstalledComponent) error {
	if _, ok := v.Components[component.Type]; ok {
		return errs.ErrComponentTypeAlreadyInstalled
	}
	v.Components[component.Type] = &component
	return nil
}
