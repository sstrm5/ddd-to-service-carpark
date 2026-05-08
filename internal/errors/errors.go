package errors

import "errors"

var (
	// domain errors
	ErrInvalidVehicleMake            = errors.New("некорректная марка автомобиля")
	ErrInvalidVehicleModel           = errors.New("некорректная модель автомобиля")
	ErrInvalidVehicleID              = errors.New("некорректный ID автомобиля")
	ErrInvalidDate                   = errors.New("некорректная дата")
	ErrVehicleUnderRepair            = errors.New("автомобиль уже в сервисе")
	ErrVehicleWritenOff              = errors.New("автомобиль списан")
	ErrInvalidVehicleStatus          = errors.New("неизвестный статус автомобиля")
	ErrInvalidOrderStatus            = errors.New("некорректный статус заказа")
	ErrMismatchTypeOrderLine         = errors.New("некорректный тип линии заказа")
	ErrInvalidWorkType               = errors.New("некорректный тип работы")
	ErrOrderLineAlreadyExists        = errors.New("такой тип строчки в заказе уже существует")
	ErrInvalidComponent              = errors.New("некорректная деталь")
	ErrComponentTypeAlreadyInstalled = errors.New("компонент такого типа уже установлен")

	// valueobjects errors
	ErrInvalidVinNumberLength     = errors.New("некорректная длина VIN номера")
	ErrInvalidVinNumberCharacters = errors.New("некорректные символы в VIN номере")
	ErrInvalidMeasure             = errors.New("некорректная мера расстояния")
	ErrInvalidDistance            = errors.New("некорректная дистанция пробега")
	ErrInvalidStatus              = errors.New("некорректный статус")
	ErrInvalidYear                = errors.New("некорректный год выпуска")
	ErrInvalidComponentType       = errors.New("некорректный тип детали")
	ErrInvalidMonthResource       = errors.New("некорректный срок службы (в месяцах)")
	ErrInvalidCurrency            = errors.New("некорректная валюта")
	ErrInvalidDescriptionLength   = errors.New("некорректная длина описания")
	ErrMileageCannotDecrease      = errors.New("пробег не может быть меньше, чем в прошлое ТО")
)
