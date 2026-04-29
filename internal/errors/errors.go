package errors

import "errors"

var (
	// domain errors
	ErrInvalidVehicleMake  = errors.New("некорректная марка автомобиля")
	ErrInvalidVehicleModel = errors.New("некорректная модель автомобиля")

	// valueobjects errors
	ErrInvalidVinNumberLength     = errors.New("некорректная длина VIN номера")
	ErrInvalidVinNumberCharacters = errors.New("некорректные символы в VIN номере")
	ErrInvalidMeasure             = errors.New("некорректная мера расстояния")
	ErrInvalidDistance            = errors.New("некорректная дистанция пробега")
	ErrInvalidStatus              = errors.New("некорректный статус")
	ErrInvalidYear                = errors.New("некорректный год выпуска")
	ErrInvalidComponentType       = errors.New("некорректный тип детали")
	ErrInvalidMonthResource       = errors.New("некорректный срок службы (в месяцах)")
)
