package entity

import (
	errs "dddcarpark/internal/errors"
	vo "dddcarpark/internal/model/valueobject"
)

type Vehicle struct {
	Vin     vo.VinNumber
	Make    string
	Model   string
	Mileage vo.Mileage
	Year    vo.Year
	Status  vo.Status
}

func NewCreate(
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
